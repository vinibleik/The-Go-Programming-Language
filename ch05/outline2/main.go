package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatalf("Too few command-line arguments: %d", len(os.Args[1:]))
	}

	for _, url := range os.Args[1:] {
		doc, err := parseHTML(url)
		if err != nil {
			log.Fatalf("outline2: %v", err)
		}

		forEachNode(doc, startElement, endElement)
	}
}

// parseHTML try to parse an html file
func parseHTML(pathFile string) (*html.Node, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}

	root, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	return root, nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, pos func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Printf("c: %v\n", c)
		forEachNode(c, pre, pos)
	}

	if pos != nil {
		pos(n)
	}
}

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}
}
