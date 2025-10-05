package commands

import (
	"fmt"
	"os/exec"
)

func Docker(args []string) error {

	command := args[0]

	switch command {

	case "dps":
		return executeDockerCommand("ps")

	case "du":
		return executeDockerCommand("system", "df")

	case "dpn":
		if len(args) < 2 {
			return executeDockerCommand("system", "prune", "-a")
		}
		switch args[1] {
		case "-i":
			return executeDockerCommand("image", "prune", "-a")
		case "-c":
			return executeDockerCommand("container", "prune")
		case "-v":
			return executeDockerCommand("volume", "prune")
		case "-n":
			return executeDockerCommand("network", "prune")
		default:
			return fmt.Errorf("error: unknown prune option '%s'", args[1])
		}

	default:
		helpText := `Promptly Docker commands:
						dps         List running containers.
						du          Show Docker disk usage.
						dpn         Prune unused Docker resources (containers, networks, images, volumes).
						dpn -i      Prune unused Docker images.
						dpn -c      Prune stopped Docker containers.
						dpn -v      Prune unused Docker volumes.
						dpn -n      Prune unused Docker networks.
						Usage:
						pmt <command> [arguments]
					`
		fmt.Println(helpText)
		return nil
	}
}

func executeDockerCommand(command string, args ...string) error {
	cmd := exec.Command("docker", append([]string{command}, args...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing docker %s: %v\nOutput: %s", command, err, string(output))
	}
	fmt.Println(string(output))
	return nil
}
