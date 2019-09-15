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

// Option is a variadic configuration option for Smallsh.
type Option func(*Smallsh)

// withStdout changes the stdout of Smallsh to the given file.
//
// For testing only.
func withStdout(out *os.File) Option {
	return func(s *Smallsh) {
		s.outFile = out
	}
}

// withStdin changes the stdin of Smallsh to the given file.
//
// For testing only.
func withStdin(in *os.File) Option {
	return func(s *Smallsh) {
		s.inFile = in
	}
}

// Smallsh handles the state of the shell.
type Smallsh struct {
	lastExitType exitType
	lastCode     int
	outFile      *os.File
	inFile       *os.File
}

// NewSmallsh creates a new Smallsh with the given options.
func NewSmallsh(opts ...Option) *Smallsh {
	s := &Smallsh{
		outFile: os.Stdout,
		inFile:  os.Stdin,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Run starts a smallsh terminal.
func (s *Smallsh) Run() error {
	stdin, stdout := os.Stdin, os.Stdout
	os.Stdin, os.Stdout = s.inFile, s.outFile
	defer func() {
		os.Stdin, os.Stdout = stdin, stdout
	}()

	for {
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("ReadString error: %v", err)
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
