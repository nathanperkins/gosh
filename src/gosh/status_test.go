package gosh

import (
	"testing"
)

func TestStatusWithExit(t *testing.T) {
	g, err := newTestGosh()
	if err != nil {
		t.Fatalf("Could not setup gosh: %v", err)
	}

	got := g.RunTest("status")
	want := `: exit value 0
: `
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

	got := g.RunTest("status")
	want := `: terminated by signal 2
: `
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
