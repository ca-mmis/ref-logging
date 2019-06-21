package main

import (
	"github.com/ca-mmis/ref-logging/projects/go/logging/logcli"
)

func exec(args []string) int {

	status := 0
	if len(args) == 1 {
		logcli.Info()
	}
	if len(args) > 1 {
		logcli.Init()
		logcli.Info()
		status = logcli.Exec(args)
	}
	return status
}

func Status() string {
	return "OKAY"
}
