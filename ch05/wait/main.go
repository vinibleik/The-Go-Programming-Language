package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	readRunes()
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err := WaitForServer(os.Args[1]) ;err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)
		log.Fatalf("Site is down: %v\n", err)
	}

}

func readRunes() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, b, err := in.ReadRune()
		if err == io.EOF {
			fmt.Printf("EOF Error: %v\n", err)
			return
		}

		if err != nil {
			log.Fatalf("read failed: r=%v, b=%v, err=%v", r, b, err)
		}

		fmt.Printf("r=%v, b=%v, err=%v\n", r, b, err)
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}