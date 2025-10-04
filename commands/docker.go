package commands

import (
	"fmt"
)

func Docker(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("error: please provide a Docker command (or run 'pmt docker help' for more information)")
	}

	command := args[0]
	commandArgs := args[1:]

	switch command {

	case "dusg": // Docker usage
		return nil // implement later

	case "dcln": // Docker clean
		return nil // implement later

	default:
		helpText := `Promptly Docker commands:
						
						Usage:
						pmt <command> [arguments]
					`
		fmt.Println(helpText)
		return nil
	}
}
