package gosh

import (
	"fmt"
	"os"
)

func (Gosh) cd(args []string) error {
	if len(args) == 0 {
		home, ok := os.LookupEnv("HOME")
		if !ok {
			return fmt.Errorf("could not find $HOME")
		}
		if err := os.Chdir(home); err != nil {
			return err
		}
	} else {
		if err := os.Chdir(args[0]); err != nil {
			return err
		}
	}
	return nil
}
