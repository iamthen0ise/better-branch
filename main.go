package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var RE = regexp.MustCompile(JIRA_RE)

const (
	// ISSUE FLAVOUR REGEX
	JIRA_RE = `([A-Z]+-[\d]+)`

	// ANSI TERMINAL CURSOR MANIPULATION
	RESET             = "\u001b[0m"
	SAVE_POSITION     = "\033[s"
	RESTORE_POSITION  = "\033[u"
	ERASE_LINE_TO_END = "\033[K"

	// ANSI TERMINAL FOREGROUND COLOR CODES
	MAGENTA      = "\u001b[36m"
	YELLOW       = "\u001b[33m"
	BRIGHT_BLACK = "\u001b[30;1m"
)

func colorize(t *string, c string) string {
	var sb strings.Builder

	sb.WriteString(c)
	sb.WriteString(*t)
	sb.WriteString(RESET)

	return sb.String()

}

type BranchArgs struct {
	issueID    string
	customText string
	value      string
}

func parseIssueId(t string) string {
	var matches []string = RE.FindAllString(t, -1)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func (b *BranchArgs) updateValue() {
	var sb strings.Builder

	sb.WriteString(b.issueID)
	if len(b.issueID) > 0 && len(b.customText) > 0 {
		sb.WriteString("-")
	}
	sb.WriteString(strings.Join(strings.Fields(b.customText), "-"))

	b.value = sb.String()
}

func (b *BranchArgs) updateFields(i string, t string) {
	b.issueID = i
	b.customText = t
	b.updateValue()
}

func (b *BranchArgs) createBranch(checkout bool) error {
	err := executeGitCommand(b.value, checkout)
	return err
}

func executeGitCommand(branchName string, checkout bool) error {
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

func main() {
	var branchArgs = BranchArgs{}

	var issueID string
	var customText string

	var withCheckout bool = true

	flag.StringVar(&issueID, "i", "", "JIRA Link or issue")
	flag.StringVar(&customText, "t", "", "Custom Issue Text")
	flag.BoolVar(&withCheckout, "c", true, "Checkout to new branch (default `true`")

	flag.Parse()

	if len(os.Args[1:]) < 1 {
		scanBranch := bufio.NewScanner(os.Stdin)

		var promptDefault string = "Enter JIRA Issue link or some text. \nPress Enter once to submit or twice when ready > "
		var promptCurrentLine *string = &branchArgs.value

		fmt.Print(SAVE_POSITION)

		for {
			fmt.Print(RESTORE_POSITION, ERASE_LINE_TO_END)

			if len(*promptCurrentLine) > 0 {
				fmt.Print("Branch Name is: ", colorize(promptCurrentLine, YELLOW), " [Enter to finish] > ")
			} else {
				fmt.Print(colorize(&promptDefault, YELLOW), " ")
			}

			scanBranch.Scan()
			text := scanBranch.Text()
			matchedText := parseIssueId(text)

			if matchedText != "" {
				issueID = matchedText
			} else if len(text) > 0 {
				customText = text
			} else {
				break
			}

			branchArgs.updateFields(issueID, customText)
		}
	} else {
		matchedText := parseIssueId(flag.Arg(0))

		if matchedText != "" {
			customText := strings.Join(flag.Args()[1:], "-")
			branchArgs.updateFields(matchedText, customText)
		} else {
			branchArgs.updateFields(branchArgs.issueID, customText)
		}

	}

	fmt.Println("Your branch name is:", colorize(&branchArgs.value, MAGENTA), "Do you want to continue and create branch? [y/N]")

	scannerCreateBranch := bufio.NewScanner(os.Stdin)
	scannerCreateBranch.Scan()

	switch strings.ToLower(scannerCreateBranch.Text()) {
	case "y":
		err := branchArgs.createBranch(withCheckout)
		if err != nil {
			fmt.Print("Something went wrong,", err.Error())
		}
	}
}
