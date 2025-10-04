package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func Git(args []string) error {

	command := args[0]
	commandArgs := args[1:]

	switch command {

	case "gpn":
		branchNameBytes, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
		if err != nil {
			return fmt.Errorf("error: failed to resolve current branch: %v", err)
		}
		branch := strings.TrimSpace(string(branchNameBytes))
		if branch == "" {
			return fmt.Errorf("error: failed to resolve current branch: empty branch name")
		}
		return executeGitCommand("push", "--set-upstream", "origin", branch)

	case "grcmt":
		cmds := [][]string{
			{"add", "-A"},
			{"commit", "--amend", "--no-edit"},
			{"push", "--force"},
		}
		for _, cmd := range cmds {
			if err := executeGitCommand(cmd[0], cmd[1:]...); err != nil {
				return err
			}
		}
		return nil

	case "gsts":
		if len(commandArgs) < 1 {
			return executeGitCommand("stash", "-u")
		}

		if len(commandArgs) >= 2 {
			return fmt.Errorf("error: too many arguments for 'git stash' command")
		}

		return executeGitCommand("stash", "push", "-u", "-m", commandArgs[0])

	default:
		helpText := `Promptly Git commands:
						gpn           Push the current branch to origin, setting upstream if needed.
						grcmt         Amend the last commit with all current changes and force-push.
						gsts [msg]    Stash all changes (including untracked). Optionally provide a message.
						Usage:
						pmt <command> [arguments]
					`
		fmt.Println(helpText)
		return nil
	}
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
