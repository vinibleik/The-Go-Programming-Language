// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func normalizeURL(argURL string) string {
	if !strings.HasPrefix(argURL, "http://") && !strings.HasPrefix(argURL, "https://") {
		argURL = "https://" + argURL
	}

	return url.QueryEscape(argURL)
}

func fetch(argURL string, ch chan<- string) {
	argURL = normalizeURL(argURL)
	start := time.Now()
	resp, err := http.Get(argURL)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", argURL, err) // send to channel
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\tbytes %s", secs, nbytes, argURL)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
