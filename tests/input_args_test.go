package tests

import (
	"strings"
	"testing"

	p "github.io/iamthen0ise/bb/src/parsing"
)

var (
	rawParams     = []string{"f", "https://jira.atlassian.com/browse/JIRA-123", "fix", "this", "m"}
	flaggedParams = []string{"-f", "--i", "https://jira.atlassian.com/browse/JIRA-123", "--t", "-m", "fix", "this"}
)

func TestInputRawPrefix(t *testing.T) {
	var inputArgs = p.InputArgs{}
	var want string
	inputArgs.ParseArgs([]string{"f", "branch", "name"})
	want = "feature"
	if inputArgs.Prefix != want {
		t.Errorf("want %v got %v", want, inputArgs.Prefix)
	}

	inputArgs.ParseArgs([]string{"h", "branch", "name"})
	want = "hotfix"
	if inputArgs.Prefix != want {
		t.Errorf("want %v got %v", want, inputArgs.Prefix)
	}
}
func TestInputRawIssueID(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(rawParams)
	want := "JIRA-123"
	if inputArgs.IssueID != want {
		t.Errorf("want %v got %v", want, inputArgs.IssueID)
	}
}

func TestInputRawRenameStrategy(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(rawParams)
	want := "Rename"
	if inputArgs.Strategy != want {
		t.Errorf("want %v got %v", want, inputArgs.IssueID)
	}
}

func TestInputRawCustomTextParts(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(rawParams)
	if len(inputArgs.CustomTextParts) < 1 {
		t.Error("customTextParts is empty")
	}
}

func TestInputFlaggedArgs(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(flaggedParams)

	want := "feature"
	if inputArgs.Prefix != want {
		t.Errorf("want %v got %v", want, inputArgs.Prefix)
	}
}
func TestInputFlaggedIssueID(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(flaggedParams)
	want := "JIRA-123"
	if inputArgs.IssueID != want {
		t.Errorf("want %v got %v", want, inputArgs.IssueID)
	}
}

func TestInputRenameStrategyArgs(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(flaggedParams)
	want := "Rename"
	if inputArgs.Strategy != want {
		t.Errorf("want %v got %v", want, inputArgs.Strategy)
	}
}

func TestInputFlaggedCustomTextParts(t *testing.T) {
	var inputArgs = p.InputArgs{}
	inputArgs.ParseArgs(flaggedParams)
	if len(inputArgs.CustomTextParts) < 1 {
		t.Error("customTextParts is empty")
	}

	joinedText := strings.Join(inputArgs.CustomTextParts, "-")
	want := "fix-this"
	if joinedText != want {
		t.Errorf("want %v got %v", "fix-this", joinedText)
	}
}

func TestInputMixedArgs(t *testing.T) {
	var params = []string{"-f", "--i", "https://jira.atlassian.com/browse/JIRA-123", "fix", "this"}
	var inputArgs = p.InputArgs{}

	inputArgs.ParseArgs(params)

	if inputArgs.Prefix != "feature" {
		t.Errorf("want %v got %v", "feature", inputArgs.Prefix)
	}
	if inputArgs.IssueID != "JIRA-123" {
		t.Errorf("want %v got %v", "JIRA-123", inputArgs.IssueID)
	}
	if len(inputArgs.CustomTextParts) < 1 {
		t.Error("customTextParts is empty")
	}

	var joinedText = strings.Join(inputArgs.CustomTextParts, "-")
	if joinedText != "fix-this" {
		t.Errorf("want %v got %v", "fix-this", joinedText)
	}
}
