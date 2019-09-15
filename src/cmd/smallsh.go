package main

import (
	"github.com/nathanperkins/smallsh-go/src/smallsh"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := smallsh.NewSmallsh()

	if err := s.Run(); err != nil {
		log.Errorf("Smallsh error: %v", err)
	}
}
