package gosh

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Exec executes the given command
func (g *Gosh) Exec(cmd string, argv []string) {
	attr := &os.ProcAttr{
		Files: []*os.File{
			g.inFile,
			g.outFile,
			g.errFile,
		},
	}
	cmd, _ = exec.LookPath(cmd)
	proc, err := os.StartProcess(cmd, argv, attr)
	if err != nil {
		fmt.Fprintf(g.errFile, "%s: %s\n", cmd, err)
		return
	}
	stat, err := proc.Wait()
	if err != nil {
		fmt.Fprintf(g.errFile, "Could not wait for %s\n", cmd)
		return
	}
	g.lastWaitStatus = stat.Sys().(syscall.WaitStatus)
}
