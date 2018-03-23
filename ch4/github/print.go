package github

import (
//	"strings"
	"fmt"
)

func PrintIssue(item *Issue, verbose bool) {
	fmt.Printf("#%-5d %9.9s %s\n",
		item.Number, item.User.Login, item.Title)
	if verbose {
		fmt.Printf("\t%.20v  -   %s\n", item.CreatedAt, item.HTMLURL)
	}
}

func PrintIssues(items []*Issue, verbose bool) {
	for _, item := range items {
		PrintIssue(item, verbose)
	}
}
