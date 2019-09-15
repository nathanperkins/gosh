package main_test

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

var (
	goshPath = flag.String("gosh", "cmd", "Path to the gosh binary.")
)

type procPipes struct {
	inWriter, inReader   *os.File
	outReader, outWriter *os.File
	errReader, errWriter *os.File
}

func newProcPipes() (*procPipes, error) {
	pipes := new(procPipes)

	var err error
	pipes.inReader, pipes.inWriter, err = os.Pipe()
	if err != nil {
		return nil, err
	}
	pipes.outReader, pipes.outWriter, err = os.Pipe()
	if err != nil {
		return nil, err
	}
	pipes.errReader, pipes.errWriter, err = os.Pipe()
	if err != nil {
		return nil, err
	}
	return pipes, nil
}

func newGoshProc(name string) (*os.Process, *procPipes, error) {
	pipes, err := newProcPipes()
	if err != nil {
		return nil, nil, err
	}
	procAttr := &os.ProcAttr{
		Files: []*os.File{
			pipes.inReader,
			pipes.outWriter,
			pipes.errWriter,
		},
	}
	proc, err := os.StartProcess(name, []string{name}, procAttr)
	if err != nil {
		return nil, nil, err
	}
	return proc, pipes, nil
}

func TestExit(t *testing.T) {
	path, _ := bazel.Runfile(*goshPath)

	proc, pipes, err := newGoshProc(path)
	if err != nil {
		t.Fatalf("Could not start cmd: %v", err)
	}
	if _, err := pipes.inWriter.WriteString("exit\n"); err != nil {
		t.Fatalf("Could not write to stdin: %v", err)
	}
	done := make(chan error)
	go func() {
		_, err := proc.Wait()
		done <- err
	}()
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("proc.Wait() err: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Timed out waiting for process.")
	}
}
