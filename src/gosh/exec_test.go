package gosh

import (
	"syscall"
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name           string
		cmd            []string
		wantStdout     string
		wantStderr     string
		wantWaitStatus syscall.WaitStatus
	}{
		{
			name:           "echo_hi",
			cmd:            []string{"echo", "hi"},
			wantStdout:     "hi\n",
			wantWaitStatus: 0x00,
		},
		{
			name:           "cat_example_txt",
			cmd:            []string{"cat", "testdata/example.txt"},
			wantStdout:     "just a few words\n",
			wantWaitStatus: 0x00,
		},
		{
			name:           "cat_invalid_file",
			cmd:            []string{"cat", "badfile.txt"},
			wantStderr:     "cat: badfile.txt: No such file or directory\n",
			wantWaitStatus: 0x100,
		},
		{
			name:           "bad_exec",
			cmd:            []string{"badfile"},
			wantStderr:     "badfile: no such file or directory\n",
			wantWaitStatus: 0x100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g, _ := newTestGosh()
			g.Exec(test.cmd[0], test.cmd)
			stdout, stderr := g.getOutput()
			if stdout != test.wantStdout {
				t.Errorf("stdout = %q, want %q", stdout, test.wantStdout)
			}
			if stderr != test.wantStderr {
				t.Errorf("stderr = %q, want %q", stderr, test.wantStderr)
			}
			if g.lastWaitStatus != test.wantWaitStatus {
				t.Errorf("status = %#x, want %#x", g.lastWaitStatus, test.wantWaitStatus)
			}
		})
	}
}
