package main

import (
	"github.com/nathanperkins/gosh/src/gosh"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := gosh.NewGosh()

	if err := s.Run(); err != nil {
		log.Errorf("Gosh error: %v", err)
	}
}
