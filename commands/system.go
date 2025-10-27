package commands

import (
	"fmt"
)

func System(args []string) error {

	command := args[0]
	// commandArgs := args[1:]

	switch command {

	default:
		helpText := `Scriptly System commands:
						sysinfo       Display system information.
						Usage:
						sc <command> [arguments]
					`
		fmt.Println(helpText)
		return nil
	}
}
