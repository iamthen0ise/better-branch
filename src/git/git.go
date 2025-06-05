package git

import (
	"os/exec"
	"strings"
)

func CreateNewBranch(branchName string, checkout bool) error {
	app := "git"

	var subcommand string
	if checkout {
		subcommand = "checkout"
	} else {
		subcommand = "branch"
	}

	cmdArgs := []string{subcommand}
	if checkout {
		cmdArgs = append(cmdArgs, "-b")
	}
	cmdArgs = append(cmdArgs, branchName)

	cmd := exec.Command(app, cmdArgs...)
	_, err := cmd.Output()

	return err
}

func RenameCurrentBranch(newBranchName string) error {
	cmd := exec.Command("git", "branch", "-m", newBranchName)
	_, err := cmd.Output()

	return err
}

func CommitChanges(messagePrefix string, messageText string) error {
	cmd := exec.Command("git", "commit", "-m", strings.Join([]string{messagePrefix, messageText}, "."))
	_, err := cmd.Output()

	return err
}
