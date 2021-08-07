package tests

import (
	"testing"

	p "github.io/iamthen0ise/bb/src/parsing"
)

func TestBuildBranchName(t *testing.T) {
	var gitBranchName = p.GitBranchName{}

	gitBranchName.IssueID = "JIRA-123"
	gitBranchName.BuildBranchName()

	want := "jira-123"
	if gitBranchName.BranchName != want {
		t.Errorf("want %v got %v", want, gitBranchName.BranchName)
	}

	gitBranchName.Prefix = "feature"
	gitBranchName.BuildBranchName()
	want = "feature/jira-123"
	if gitBranchName.BranchName != want {
		t.Errorf("want %v got %v", want, gitBranchName.BranchName)
	}

	gitBranchName.CustomTextParts = append(gitBranchName.CustomTextParts, "Some", "text", "i", "want", "to", "add")
	gitBranchName.BuildBranchName()
	want = "feature/jira-123-some-text-i-want-to-add"
	if gitBranchName.BranchName != want {
		t.Errorf("want %v got %v", want, gitBranchName.BranchName)
	}
}
