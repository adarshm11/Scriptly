package commands

import (
	"fmt"
)

func Help() error {
	helpText := `Promptly Commands:
		pmt system              System-related commands (e.g., storage checks).
		pmt docker             Docker-related commands (e.g., view containers, clean up).
		pmt git                 Git-related commands (e.g., pushnew, recommit, stash).
		Usage:
		pmt <command> [arguments]
		For more information on a specific command, run 'pmt <command> help'.
	`
	fmt.Println(helpText)
	return nil
}
