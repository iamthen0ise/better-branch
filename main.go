package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	p "github.io/iamthen0ise/bb/src/parsing"
	s "github.io/iamthen0ise/bb/src/screen"
)

func main() {
	var (
		inputArgs     = p.InputArgs{}
		gitBranchName = p.GitBranchName{}
	)

	flag.String("i", "", "JIRA Link or issue")
	flag.String("t", "", "Custom Issue Text")
	flag.Bool("f", false, "Set `feature` prefix")
	flag.Bool("h", false, "Set `hotfix` prefix")
	flag.Bool("c", true, "Checkout to new branch (default `true`")
	flag.Parse()

	if len(os.Args) > 1 {
		inputArgs.ParseArgs(os.Args[1:])
		gitBranchName.UpdateFields(inputArgs.Prefix, inputArgs.IssueID, inputArgs.CustomTextParts)

	} else {
		inputScanner := bufio.NewScanner(os.Stdin)

		var promptDefault = "Enter JIRA Issue link or some text. Press Enter once to submit or twice when ready > "
		var promptCurrentLine = &gitBranchName.BranchName

		fmt.Print(s.SavePosition)

		for {
			fmt.Print(s.RestorePosition, s.EraseLineToEnd)

			if len(*promptCurrentLine) > 0 {
				fmt.Print("Branch Name is: ", s.Colorize(promptCurrentLine, s.Yellow), " [Enter to finish] > ")
			} else {
				fmt.Print(s.Colorize(&promptDefault, s.Yellow), " ")
			}

			inputScanner.Scan()
			text := inputScanner.Text()
			inputArgs.ParseArg(&text)

			if inputArgs.IssueID == "" && len(text) > 1 {
				inputArgs.CustomTextParts = append(inputArgs.CustomTextParts, text)
			} else if len(text) == 0 {
				break
			}
			gitBranchName.UpdateFields(inputArgs.Prefix, inputArgs.IssueID, inputArgs.CustomTextParts)
		}
	}

	fmt.Println("Your branch name is:", s.Colorize(&gitBranchName.BranchName, s.Magenta))
	fmt.Println("Do you want to continue and create branch? [Enter to continue]")

	scannerCreateBranch := bufio.NewScanner(os.Stdin)
	scannerCreateBranch.Scan()

	switch strings.ToLower(scannerCreateBranch.Text()) {
	case "":
		err := gitBranchName.CreateBranch(true)
		if err != nil {
			fmt.Print("Something went wrong,", err.Error())
		}
	}
}
