package main

import (
	"github.com/rmrfslashbin/ssmparams/pkg/cmds"
	"github.com/sirupsen/logrus"
)

// main is the entry point for the application.
func main() {
	// Catch errors
	var err error
	defer func() {
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("main crashed")
		}
	}()

	// Run the CLI
	cmds.Execute()
}
