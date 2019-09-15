package gosh

import (
	"testing"
)

func TestStatusWithExit(t *testing.T) {
	g, err := newTestGosh()
	if err != nil {
		t.Fatalf("Could not setup gosh: %v", err)
	}

	g.Status()
	got, _ := g.getOutput()
	want := "exit value 0\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStatusWithSignal(t *testing.T) {
	g, err := newTestGosh()
	if err != nil {
		t.Fatalf("Could not setup gosh: %v", err)
	}
	g.Gosh.lastExitType = exitTypeSignal
	g.Gosh.lastCode = 2

	g.Status()
	got, _ := g.getOutput()
	want := "terminated by signal 2\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
