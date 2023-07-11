package main

import (
	"ch04/text-and-html/github"
	"fmt"
)

func withoutTemplate(result *github.IssueSearchResult) {
	fmt.Println("WITHOUT TEMPLATE!")
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, issue := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
}