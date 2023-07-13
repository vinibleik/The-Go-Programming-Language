package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Printf("\nURL: %s\n", url)
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex5.5: %v\n", err)
			continue
		}

		fmt.Printf("Words: %d Images: %d\n", words, images)
	}
}

// Wr ite a program wordfreq to rep ort the fre quency of each word in an inp ut text
// file. Cal l input.Split(bufio.ScanWords) before the firs t call to Scan to bre ak the inp ut int o
// word s instead of lines.

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return 
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
	}

	wordsFirstChild, imgFirstChild := countWordsAndImages(n.FirstChild)
	wordsNextSibling, imgNextSibling := countWordsAndImages(n.NextSibling)

	words += wordsFirstChild + wordsNextSibling
	images += imgFirstChild + imgNextSibling

	return
}