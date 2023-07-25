// Prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	HTTP  = "http://"
	HTTPS = "https://"
)

func normalizeURL(url string) string {
	if strings.HasPrefix(url, HTTP) || strings.HasPrefix(url, HTTPS) {
		return url
	}

	return HTTPS + url
}

func fetchURL(url string) (int64, error) {
	url = normalizeURL(url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return io.Copy(os.Stdout, resp.Body)
}

func main() {
	for _, url := range os.Args[1:] {
		_, err := fetchURL(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1.7: %v\n", err)
			continue
		}
	}
}
