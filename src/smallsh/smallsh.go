package smallsh

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	prompt = ": "
)

type exitType int

const (
	exitTypeExit exitType = iota
	exitTypeSignal
)

type Smallsh struct {
	lastExitType exitType
	lastCode     int
}

func (s *Smallsh) Run() error {
	for {
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		if input == "" || strings.HasPrefix(input, "#") {
			continue
		}
		if input == "exit" {
			return nil
		}
		inputSplit := strings.Split(input, " ")
		if inputSplit[0] == "cd" {
			if err := s.cd(inputSplit[1:]); err != nil {
				log.Error(err)
			}
		} else {
			fmt.Printf("Not implemented: %v\n", inputSplit)
		}
	}
}
