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

	case "gi":
		return executeGitCommand("init")

	case "gcl":
		return executeGitCommand("clone", commandArgs...)

	case "gst":
		return executeGitCommand("status")

	case "ga":
		if len(commandArgs) < 1 {
			return executeGitCommand("add", "-A")
		}
		return executeGitCommand("add", commandArgs...)

	case "gc":
		if len(commandArgs) < 1 {
			return fmt.Errorf("error: commit message is required for 'git commit'")
		}
		return executeGitCommand("commit", "-m", strings.Join(commandArgs, " "))

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

	case "gp":
		return executeGitCommand("push")

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

	case "gpl":
		return executeGitCommand("pull")

	case "gf":
		return executeGitCommand("fetch")

	case "gsts":
		if len(commandArgs) < 1 {
			return executeGitCommand("stash", "-u")
		}

		if len(commandArgs) >= 2 {
			switch commandArgs[0] {

			case "-l":
				return executeGitCommand("stash", "list")

			case "-d":
				if len(commandArgs) > 1 {
					return executeGitCommand("stash", "drop", commandArgs[1])
				}
				return executeGitCommand("stash", "drop")

			case "-p":
				return executeGitCommand("stash", "pop")
			}
		}

		return executeGitCommand("stash", "push", "-u", "-m", commandArgs[0])

	case "gb":
		if len(commandArgs) == 0 {
			return executeGitCommand("branch") // List branches
		}

		switch commandArgs[0] {

		case "-d", "-D":
			if len(commandArgs) > 1 {
				return executeGitCommand("branch", commandArgs[0], commandArgs[1]) // Delete branch
			} else {
				return fmt.Errorf("error: branch name is required for 'git branch %s'", commandArgs[0])
			}

		case "-r":
			if len(commandArgs) > 1 {
				return executeGitCommand("push", "origin", "--delete", commandArgs[1]) // Delete remote branch
			} else {
				return fmt.Errorf("error: remote branch name is required for 'git branch -r'")
			}

		default:
			return executeGitCommand("branch", commandArgs[0]) // Create a new branch
		}

	case "gchk":
		if len(commandArgs) < 1 {
			return fmt.Errorf("error: branch name is required for 'git checkout'")
		}

		if commandArgs[0] == "-b" {
			if len(commandArgs) < 2 {
				return fmt.Errorf("error: new branch name is required for 'git checkout -b'")
			}
			return executeGitCommand("checkout", "-b", commandArgs[1]) // Create and checkout new branch
		}
		return executeGitCommand("checkout", commandArgs[0]) // Checkout existing branch

	case "grb":
		if len(commandArgs) < 1 {
			return fmt.Errorf("error: branch name is required for 'git rebase'")
		}
		return executeGitCommand("rebase", commandArgs[0])

	case "gmg":
		if len(commandArgs) < 1 {
			return fmt.Errorf("error: tag name is required for 'git merge'")
		}
		return executeGitCommand("merge", commandArgs[0])

	default:
		helpText := `
						gi            Initialize a new Git repository.
						gcl <url>     Clone a Git repository from the provided URL.
						gst           Show the working tree status.
						ga [file...]  Stage all changes, or specific files if provided.
						gc <message>  Commit staged changes with the provided commit message.
						gp            Push the current branch to origin.
						gpn           Push the current branch to origin, setting upstream if needed.
						gpl           Pull changes from the remote repository.
						gf            Fetch changes from the remote repository.
						grcmt         Amend the last commit with all current changes and force-push.
						gsts [msg]    Stash all changes (including untracked). Optionally provide a message.
						              -l: List stashes, -d [id]: Drop stash, -p: Pop stash.
						gb [branch]   List branches, create new branch, or manage branches.
						              -d/-D <branch>: Delete branch, -r <branch>: Delete remote branch.
						gchk <branch> Checkout existing branch or create new branch with -b flag.
						grb <branch>  Rebase current branch onto the specified branch.
						gmg <branch>  Merge the specified branch into the current branch.
						Usage:
						sc <command> [arguments]
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
