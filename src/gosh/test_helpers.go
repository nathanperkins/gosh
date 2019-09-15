package gosh

import (
	"io/ioutil"
	"os"
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

func (p *procPipes) Close() error {
	if err := p.inReader.Close(); err != nil {
		return err
	}
	if err := p.outReader.Close(); err != nil {
		return err
	}
	if err := p.errReader.Close(); err != nil {
		return err
	}
	if err := p.inWriter.Close(); err != nil {
		return err
	}
	if err := p.outWriter.Close(); err != nil {
		return err
	}
	if err := p.errWriter.Close(); err != nil {
		return err
	}
	return nil
}

type testGosh struct {
	*Gosh
	pipes *procPipes
}

func newTestGosh() (*testGosh, error) {
	pipes, err := newProcPipes()
	if err != nil {
		return nil, err
	}
	opts := []Option{
		withStdin(pipes.inReader),
		withStdout(pipes.outWriter),
		withStderr(pipes.errWriter),
	}
	testGosh := new(testGosh)
	testGosh.Gosh = NewGosh(opts...)
	testGosh.pipes = pipes

	return testGosh, nil
}

func (g *testGosh) RunTest(s string) string {
	defer g.pipes.Close()

	done := make(chan error)
	go func() {
		done <- g.Run()
		g.pipes.outWriter.Close()
	}()
	g.pipes.inWriter.WriteString(s + "\nexit\n")
	<-done
	out, _ := ioutil.ReadAll(g.pipes.outReader)
	return string(out)
}
