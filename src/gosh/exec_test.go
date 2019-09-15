package gosh

import "testing"

func TestExecEcho(t *testing.T) {
	g, _ := newTestGosh()
	g.Exec("echo", []string{"echo", "hi"})
	got, _ := g.getOutput()
	want := "hi\n"
	if got != want {
		t.Logf("For Exec(%q):", "echo hi")
		t.Errorf("got %q, want %q", got, want)
	}
}
