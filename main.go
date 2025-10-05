package promptly

import (
	"fmt"
	"promptly/commands"
)

func main(args []string) error {
	if len(args) <= 1 {
		// print usage information
		return fmt.Errorf("no command provided")
	}

	switch args[1][0] {

	case 'g':
		return commands.Git(args)

	case 'd':
		return commands.Docker(args)

	case 's':
		return commands.System(args)

	default:
		return commands.Help()
	}
}
