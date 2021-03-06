package parsing

import (
	"regexp"
	"strings"
)

const JiraRe = `([A-Z]+-[\d]+)`

var (
	RE            = regexp.MustCompile(JiraRe)
	FlagConstants = []string{"-i", "--i", "-t", "--t", "-c", "--c"}
)

type InputArgs struct {
	Prefix          string
	IssueID         string
	CustomTextParts []string
	Strategy        string
	ForceCreate     bool
}

func (a *InputArgs) ParseArg(t *string) {
	for _, char := range FlagConstants {
		if *t == char {
			return
		}
	}

	var issuerIdMatches = RE.FindAllString(strings.ToUpper(*t), -1)

	if len(issuerIdMatches) > 0 {
		a.IssueID = issuerIdMatches[0]
	} else if strings.Trim(*t, "-") == "f" {
		a.Prefix = "feature"
	} else if strings.Trim(*t, "-") == "h" {
		a.Prefix = "hotfix"
	} else if strings.Trim(*t, "-") == "b" {
		a.Prefix = "bugfix"
	} else if strings.Trim(*t, "-") == "r" {
		a.Prefix = "release"
	} else if strings.Trim(*t, "-") == "m" {
		a.Strategy = "Rename"
	} else if strings.Trim(*t, "-") == "y" {
		a.ForceCreate = true
	} else {
		a.CustomTextParts = append(a.CustomTextParts, *t)
	}
}

func (a *InputArgs) ParseArgs(s []string) {
	for _, c := range s {
		a.ParseArg(&c)
	}
}
