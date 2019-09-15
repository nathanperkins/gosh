package main

import (
	"github.com/nathanperkins/smallsh-go/smallsh"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := new(smallsh.Smallsh)

	if err := s.Run(); err != nil {
		log.Errorf("Error: %v", err)
	}
}
