package main

import (
	"ch04/text-and-html/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args[1:]) % 2 != 0 {
		html_template(result)
	} else {
		withoutTemplate(result)
		text_template(result)
	}
}