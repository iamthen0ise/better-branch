package parsing

import (
	"regexp"
	"strings"
)

const JiraRe = `([A-Z]+-[\d]+)`
const KaitenRe = `(KAITEN.(RU|COM)\/[\d]+)`

var (
	JiraReCompiled   = regexp.MustCompile(JiraRe)
	KaitenReCompiled = regexp.MustCompile(KaitenRe)
	FlagConstants    = []string{"-i", "--i", "-t", "--t", "-c", "--c"}
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

	var jiraMatches = JiraReCompiled.FindAllString(strings.ToUpper(*t), -1)
	var kaitenMatches = KaitenReCompiled.FindAllString(strings.ToUpper(*t), -1)

	var issueIdMatches string

	if len(jiraMatches) > 0 {
		issueIdMatches = jiraMatches[0]
	} else if len(kaitenMatches) > 0 {
		issueIdMatches = strings.Replace(kaitenMatches[0], "/", "-", -1)
		issueIdMatches = strings.Replace(issueIdMatches, ".RU", "", -1)
		issueIdMatches = strings.Replace(issueIdMatches, ".COM", "", -1)
	} else {
		issueIdMatches = ""
	}

	if len(issueIdMatches) > 0 {
		a.IssueID = issueIdMatches
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
