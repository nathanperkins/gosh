package gosh

import "fmt"

func (g *Gosh) Status() {
	if g.lastExitType == exitTypeExit {
		fmt.Printf("exit value %d\n", g.lastCode)
	} else {
		fmt.Printf("terminated by signal %d\n", g.lastCode)
	}
}
