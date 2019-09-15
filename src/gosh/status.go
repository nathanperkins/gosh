package gosh

import "fmt"

// Status writes the last exit or signal to the output.
func (g *Gosh) Status() {
	if g.lastExitType == exitTypeExit {
		fmt.Fprintf(g.outFile, "exit value %d\n", g.lastCode)
	} else {
		fmt.Fprintf(g.outFile, "terminated by signal %d\n", g.lastCode)
	}
}
