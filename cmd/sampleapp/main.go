package main

import (
	"github.com/dyxj/gobase/config"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.Help()
	if err != nil {
		logrus.Fatalf("broken build: %v", err)
	}

	cfg, err := config.FromEnvVar()
	if err != nil {
		logrus.Fatalf("config.FromEnvVar: %v", err)
	}

	logrus.Println(cfg)
}
