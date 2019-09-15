package gosh

import (
	"os"
	"testing"
	"time"
)

func TestExit(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Could not open pipe: %v", err)
	}
	if _, err := w.Write([]byte("exit\n")); err != nil {
		t.Fatalf("Failed to write exit to pipe: %v", err)
	}

	done := make(chan bool)
	go func() {
		s := NewGosh(withStdin(r), withStdout(nil))
		s.Run()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Errorf("Timed out.")
	}
}
