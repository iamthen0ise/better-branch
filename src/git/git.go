package git

import (
	"os/exec"
	"strings"
)

func ExecuteGitCommand(branchName string, checkout bool) error {
	app := "git"

	var subcommand string
	var args []string

	if checkout {
		subcommand = "checkout"
		args = append(args, "-b")
	} else {
		subcommand = "branch"
	}

	cmd := exec.Command(app, subcommand, strings.Join(args, " "), branchName)
	_, err := cmd.Output()

	return err
}
