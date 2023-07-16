package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	// fmt.Printf("resp.Request.URL.Path: %v\n", resp.Request.URL.Path)
	// fmt.Printf("local: %v\n", local)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("Too few command-line arguments!")
	}

	for _, url := range os.Args[1:] {
		fetch(url)
	}
}
