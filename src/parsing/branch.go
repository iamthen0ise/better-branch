package parsing

import (
	"strings"

	"github.io/iamthen0ise/bb/src/git"
)

type GitBranchName struct {
	Prefix          string
	IssueID         string
	CustomTextParts []string
	BranchName      string
}

func (o *GitBranchName) BuildBranchName() {
	var sb strings.Builder
	if len(o.Prefix) > 0 {
		sb.WriteString(o.Prefix)
		sb.WriteString("/")
	}

	sb.WriteString(o.IssueID)
	if len(o.IssueID) > 0 && len(o.CustomTextParts) > 0 {
		sb.WriteString("-")
	}
	sb.WriteString(strings.Join(o.CustomTextParts, "-"))

	o.BranchName = sb.String()
}

func (o *GitBranchName) UpdateFields(p string, i string, tp []string) {
	o.Prefix = p
	o.IssueID = i
	o.CustomTextParts = tp
	o.BuildBranchName()
}

func (o *GitBranchName) CreateBranch(checkout bool) error {
	err := git.ExecuteGitCommand(o.BranchName, checkout)
	return err
}
