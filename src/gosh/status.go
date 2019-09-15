package gosh

import "fmt"

// Status writes the last exit or signal to the output.
func (g *Gosh) Status() {
	if g.lastWaitStatus.Exited() {
		fmt.Fprintf(g.outFile, "exit value %d\n", g.lastWaitStatus.ExitStatus())
	} else if g.lastWaitStatus.Signaled() {
		fmt.Fprintf(g.outFile, "terminated by signal %d\n", g.lastWaitStatus.Signal())
	} else {
		fmt.Fprintf(g.outFile, "not exited or signaled\n")
	}
}
