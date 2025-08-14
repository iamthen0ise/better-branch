package git

import (
	"os"
	"os/exec"
	"strings"
)

func CommitChanges(messagePrefix string, tp []string) error {
	commitMessage := strings.ToLower(strings.Join(tp, " "))
	if len(commitMessage) > 0 {
		commitMessage = strings.ToUpper(commitMessage[:1]) + commitMessage[1:]
	}

	if len(messagePrefix) > 0 {
		commitMessage = messagePrefix + " " + commitMessage
	}

	cmd := exec.Command("git", "commit", "-m", commitMessage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
