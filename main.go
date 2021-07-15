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

type ParsedArgs struct {
	issueID    string
	customText string
}

func (p *ParsedArgs) parseRawInput(t string) {
	var matches []string = RE.FindAllString(t, -1)
	if len(matches) > 0 {
		p.issueID = matches[0]
	} else {
		p.customText = t
	}
}

type BranchArgs struct {
	issueID    string
	customText string
	value      string
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

	var issue string
	var ct string

	var withCheckout bool = true

	flag.StringVar(&issue, "i", "", "JIRA Link or issue")
	flag.StringVar(&ct, "t", "", "Custom Issue Text")
	flag.BoolVar(&withCheckout, "c", true, "Checkout to new branch (default `true`")

	flag.Parse()

	var parsedArgs = ParsedArgs{}

	if len(os.Args[1:]) < 1 {
		inputScanner := bufio.NewScanner(os.Stdin)

		var promptDefault string = "Enter JIRA Issue link or some text. Press Enter once to submit or twice when ready > "
		var promptCurrentLine *string = &branchArgs.value

		fmt.Print(SAVE_POSITION)

		for {
			fmt.Print(RESTORE_POSITION, ERASE_LINE_TO_END)

			if len(*promptCurrentLine) > 0 {
				fmt.Print("Branch Name is: ", colorize(promptCurrentLine, YELLOW), " [Enter to finish] > ")
			} else {
				fmt.Print(colorize(&promptDefault, YELLOW), " ")
			}

			inputScanner.Scan()
			text := inputScanner.Text()
			parsedArgs.parseRawInput(text)

			if parsedArgs.issueID == "" && len(text) > 0 {
				parsedArgs.customText = text
			} else if len(text) == 0 {
				break
			}

			branchArgs.updateFields(parsedArgs.issueID, parsedArgs.customText)
		}
	} else {
		parsedArgs.parseRawInput(issue)

		if len(parsedArgs.issueID) < 1 {
			parsedArgs.parseRawInput(flag.Arg(0))
		}

		if len(parsedArgs.customText) < 1 {
			parsedArgs.customText = ct
		}

		if flag.NArg() > 0 {
			parsedArgs.customText = parsedArgs.customText + " " + strings.Join(flag.Args()[1:], " ")
		}

		branchArgs.updateFields(parsedArgs.issueID, parsedArgs.customText)
	}

	fmt.Println("Your branch name is:", colorize(&branchArgs.value, MAGENTA))
	fmt.Println("Do you want to continue and create branch? [Enter to continue]")

	scannerCreateBranch := bufio.NewScanner(os.Stdin)
	scannerCreateBranch.Scan()

	switch strings.ToLower(scannerCreateBranch.Text()) {
	case "":
		err := branchArgs.createBranch(withCheckout)
		if err != nil {
			fmt.Print("Something went wrong,", err.Error())
		}
	}
}
