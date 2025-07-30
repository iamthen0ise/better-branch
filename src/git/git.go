package git

import (
	"os/exec"
	"strings"
)

func CreateNewBranch(branchName string, checkout bool) error {
	app := "git"

	var args []string
	if checkout {
		args = append(args, "checkout", "-b", branchName)
	} else {
		args = append(args, "branch", branchName)
	}

	cmd := exec.Command(app, args...)
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
