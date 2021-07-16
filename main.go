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
	prefix     string
	issueID    string
	customText string
}

func (p *ParsedArgs) parseRawInput(t *string) {
	var matches []string = RE.FindAllString(strings.ToUpper(*t), -1)
	if len(matches) > 0 {
		p.issueID = matches[0]
	} else if *t == "f" {
		p.prefix = "feature"
	} else if *t == "h" {
		p.prefix = "hotfix"

	} else {
		p.customText = *t
	}
}

type BranchArgs struct {
	prefix     string
	issueID    string
	customText string
	value      string
}

func (b *BranchArgs) updateValue() {
	var sb strings.Builder
	if len(b.prefix) > 0 {
		sb.WriteString(b.prefix)
		sb.WriteString("/")
	}

	sb.WriteString(b.issueID)
	if len(b.issueID) > 0 && len(b.customText) > 0 {
		sb.WriteString("-")
	}
	sb.WriteString(strings.Join(strings.Fields(b.customText), "-"))

	b.value = sb.String()
}

func (b *BranchArgs) updateFields(p string, i string, t string) {
	b.prefix = p
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
	var (
		branchArgs = BranchArgs{}

		issue  string
		ct     string
		prefix string

		feature string
		hotfix  string

		withCheckout bool = true
	)

	flag.StringVar(&issue, "i", "", "JIRA Link or issue")
	flag.StringVar(&ct, "t", "", "Custom Issue Text")
	flag.StringVar(&feature, "f", "", "set `feature` prefix")
	flag.StringVar(&hotfix, "h", "", "set `hotfix` prefix")
	flag.BoolVar(&withCheckout, "c", true, "Checkout to new branch (default `true`")

	flag.Parse()

	var pargs = ParsedArgs{}

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
			pargs.parseRawInput(&text)

			if pargs.issueID == "" && len(text) > 1 {
				pargs.customText = text
			} else if strings.ContainsAny(text, "fh") {
				pargs.customText = ""
			} else if len(text) == 0 {
				break
			}
			branchArgs.updateFields(pargs.prefix, pargs.issueID, pargs.customText)
		}
	} else {
		pargs.parseRawInput(&issue)

		if len(pargs.issueID) < 1 {
			var arg = flag.Arg(0)
			pargs.parseRawInput(&arg)
		}

		if prefix != "" {
			pargs.prefix = prefix
		} else if feature != "" {
			pargs.prefix = "feature"
		} else if hotfix != "" {
			pargs.prefix = "hotfix"
		}

		if len(pargs.customText) < 1 {
			pargs.customText = ct
		}

		if flag.NArg() > 0 {
			pargs.customText = pargs.customText + " " + strings.Join(flag.Args()[1:], " ")
		}

		branchArgs.updateFields(pargs.prefix, pargs.issueID, pargs.customText)
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
