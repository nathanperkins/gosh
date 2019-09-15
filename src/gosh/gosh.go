package gosh

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

// Option is a variadic configuration option for Gosh.
type Option func(*Gosh)

// withStdout changes the stdout of Gosh to the given file.
//
// For testing only.
func withStdout(out *os.File) Option {
	return func(s *Gosh) {
		s.outFile = out
	}
}

// withStdin changes the stdin of Gosh to the given file.
//
// For testing only.
func withStdin(in *os.File) Option {
	return func(s *Gosh) {
		s.inFile = in
	}
}

// Gosh handles the state of the shell.
type Gosh struct {
	lastExitType exitType
	lastCode     int
	outFile      *os.File
	inFile       *os.File
}

// NewGosh creates a new Gosh with the given options.
func NewGosh(opts ...Option) *Gosh {
	s := &Gosh{
		outFile: os.Stdout,
		inFile:  os.Stdin,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Run starts a gosh terminal.
func (s *Gosh) Run() error {
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
