package main

import (
	"ch04/text-and-html/github"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").
Funcs(template.FuncMap{"daysAgo": daysAgo}).
Parse(templ))

func text_template(result *github.IssueSearchResult) {
	fmt.Println("\nWITH TEMPLATE")
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}