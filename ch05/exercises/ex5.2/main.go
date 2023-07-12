package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.2: %v\n", err)
		os.Exit(1)
	}

	domElements := make(map[string]int)
	countDomElements(domElements, doc)
	fmt.Printf("%-10s\t%10s\n", "Element", "Count")
	for k, v := range domElements {
		fmt.Printf("%-10s\t%10d\n", k, v)
	}
}

func countDomElements(domElements map[string]int, dom *html.Node) {
	if dom == nil {
		return
	}

	if dom.Type == html.ElementNode {
		domElements[dom.Data]++
	}

	countDomElements(domElements, dom.FirstChild)
	countDomElements(domElements, dom.NextSibling)
}