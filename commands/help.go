package commands

import (
	"fmt"
)

func Help() error {
	helpText := `
		pmt s             System-related commands (e.g., storage checks).
		pmt d             Docker-related commands (e.g., view containers, clean up).
		pmt g             Git-related commands (e.g., push, commit, stash).
		pmt help          Show help information.

		Usage:
		pmt <command> [arguments]
		For more information on a specific command, run 'pmt <command> help'.
	`
	fmt.Println(helpText)
	return nil
}
