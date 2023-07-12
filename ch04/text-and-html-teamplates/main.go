package main

import (
	"log"
	"os"

	"gopl.vini/ch04/github"
)

func main() {
	result, err := github.SearchIssue(os.Args[1:])
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