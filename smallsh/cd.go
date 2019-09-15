package smallsh

import (
	"fmt"
	"os"
)

func (Smallsh) cd(args []string) error {
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
