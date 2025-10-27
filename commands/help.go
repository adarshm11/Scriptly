package commands

import (
	"fmt"
)

func Help() error {
	helpText := `
		sc s             System-related commands (e.g., storage checks).
		sc d             Docker-related commands (e.g., view containers, clean up).
		sc g             Git-related commands (e.g., push, commit, stash).
		sc help          Show help information.

		Usage:
		sc <command> [arguments]
		For more information on a specific command, run 'sc <command> help'.
	`
	fmt.Println(helpText)
	return nil
}
