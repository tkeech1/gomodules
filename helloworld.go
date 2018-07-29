package gomodules

import (
	log "github.com/sirupsen/logrus"
)

func sayHello() string {
	log.Info("saying hello...")
	return "Hello, world"
}
