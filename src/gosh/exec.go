package gosh

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Exec executes the given command and updates the lastWaitStatus.
func (g *Gosh) Exec(cmd string, argv []string) {
	attr := &os.ProcAttr{
		Files: []*os.File{
			g.inFile,
			g.outFile,
			g.errFile,
		},
	}
	cmdPath, _ := exec.LookPath(cmd)
	proc, err := os.StartProcess(cmdPath, argv, attr)
	if err != nil {
		fmt.Fprintf(g.errFile, "%s: no such file or directory\n", cmd)
		g.lastWaitStatus = 0x100
		return
	}
	stat, err := proc.Wait()
	if err != nil {
		fmt.Fprintf(g.errFile, "Could not wait for %s\n", cmd)
		g.lastWaitStatus = 0x100
		return
	}
	g.lastWaitStatus = stat.Sys().(syscall.WaitStatus)
}
