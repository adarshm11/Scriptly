package main

import (
	"fmt"
	"os"
	"scriptly/commands"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		err := commands.Help()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		return
	}

	switch args[1][0] {

	case 'g':
		err := commands.Git(args)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case 'd':
		err := commands.Docker(args)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case 's':
		err := commands.System(args)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	default:
		err := commands.Help()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
