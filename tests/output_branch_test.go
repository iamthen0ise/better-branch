package tests

import (
	"testing"

	p "github.io/iamthen0ise/bb/src/parsing"
)

func TestBuildBranchName(t *testing.T) {
	var gitBranchName = p.GitBranchName{}

	gitBranchName.IssueID = "JIRA-123"
	gitBranchName.BuildBranchName()

	want := "JIRA-123"
	if gitBranchName.BranchName != want {
		t.Errorf("want %v got %v", want, gitBranchName.BranchName)
	}

	gitBranchName.Prefix = "feature"
	gitBranchName.BuildBranchName()
	want = "feature/JIRA-123"
	if gitBranchName.BranchName != want {
		t.Errorf("want %v got %v", want, gitBranchName.BranchName)
	}

	gitBranchName.CustomTextParts = append(gitBranchName.CustomTextParts, "Some", "text", "i", "want", "to", "add")
	gitBranchName.BuildBranchName()
	want = "feature/JIRA-123-Some-text-i-want-to-add"
	if gitBranchName.BranchName != want {
		t.Errorf("want %v got %v", want, gitBranchName.BranchName)
	}
}
