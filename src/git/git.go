package git

import (
	"os/exec"
	"strings"
)

func CreateNewBranch(branchName string, checkout bool) error {
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
