package gosh

import (
	"bufio"
	"fmt"
	"io"
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
	return func(g *Gosh) {
		g.outFile = out
	}
}

// withStdin changes the stdin of Gosh to the given file.
//
// For testing only.
func withStdin(in *os.File) Option {
	return func(g *Gosh) {
		g.inFile = in
	}
}

// withStderr changes the stdin of Gosh to the given file.
//
// For testing only.
func withStderr(err *os.File) Option {
	return func(g *Gosh) {
		g.errFile = err
	}
}

// Gosh handles the state of the shell.
type Gosh struct {
	lastExitType exitType
	lastCode     int
	inFile       *os.File
	outFile      *os.File
	errFile      *os.File
}

// NewGosh creates a new Gosh with the given options.
func NewGosh(opts ...Option) *Gosh {
	g := &Gosh{
		inFile:  os.Stdin,
		outFile: os.Stdout,
		errFile: os.Stderr,
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

// Run starts a gosh terminal.
func (g *Gosh) Run() error {
	stdin, stdout, stderr := os.Stdin, os.Stdout, os.Stderr
	os.Stdin, os.Stdout, os.Stderr = g.inFile, g.outFile, g.errFile
	defer func() {
		os.Stdin, os.Stdout, os.Stderr = stdin, stdout, stderr
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("ReadString error: %v", err)
		}
		input = strings.TrimSpace(input)
		if input == "" || strings.HasPrefix(input, "#") {
			continue
		}
		if input == "exit" {
			return nil
		}
		if input == "status" {
			g.Status()
			continue
		}
		inputSplit := strings.Split(input, " ")
		if inputSplit[0] == "cd" {
			if err := g.cd(inputSplit[1:]); err != nil {
				log.Error(err)
			}
		} else {
			fmt.Printf("Not implemented: %v\n", inputSplit)
		}
	}
}
