package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func Git(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("error: please provide a Git command (or run 'pmt git help' for more information)")
	}

	subcommand := args[0]
	subcommandArgs := args[1:]

	switch subcommand {

	case "pushnew":
		branchNameBytes, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
		if err != nil {
			return fmt.Errorf("error: failed to resolve current branch: %v", err)
		}
		branch := strings.TrimSpace(string(branchNameBytes))
		if branch == "" {
			return fmt.Errorf("error: failed to resolve current branch: empty branch name")
		}
		return executeGitCommand("push", "--set-upstream", "origin", branch)

	case "recommit":
		cmds := [][]string{
			{"add", "-A"},
			{"commit", "--amend", "--no-edit"},
			{"push", "--force"},
		}
		return runGitSequence(cmds)

	case "stash":
		if len(subcommandArgs) < 1 {
			return executeGitCommand("stash", "-u")
		}

		if len(subcommandArgs) >= 2 {
			return fmt.Errorf("error: too many arguments for 'git stash' subcommand")
		}

		return executeGitCommand("stash", "push", "-u", "-m", subcommandArgs[0])

	case "help":
		helpText := `Promptly Git Subcommands:
						pushnew        Push the current branch to origin, setting upstream if needed.
						recommit       Amend the last commit with all current changes and force-push.
						stash [msg]    Stash all changes (including untracked). Optionally provide a message.
						Usage:
						pmt git <subcommand> [arguments]
					`
		fmt.Println(helpText)
		return nil

	default:
		return fmt.Errorf("error: unknown Git command: %s", subcommand)
	}
}

// runGitSequence executes a list of git subcommands in order, stopping on first error.
func runGitSequence(cmds [][]string) error {
	for _, c := range cmds {
		if len(c) == 0 {
			continue
		}
		if err := executeGitCommand(c[0], c[1:]...); err != nil {
			return err
		}
	}
	return nil
}

func executeGitCommand(command string, args ...string) error {
	cmd := exec.Command("git", append([]string{command}, args...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing git %s: %v\nOutput: %s", command, err, string(output))
	}
	fmt.Println(string(output))
	return nil
}
