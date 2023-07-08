package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	SHA256 = 256
	SHA384 = 384
	SHA512 = 512
)

func main() {

	if len(os.Args[1:]) < 1 {
		fmt.Fprint(os.Stderr, "Insufficient command-line arguments!\n")
		os.Exit(1)
	}

	var sha = flag.Int("sha", SHA256, "Number of SHA to use")
	var text = flag.String("input", "", "Text to calculate the sum")
	flag.Parse()

	if *text == "" {
		fmt.Fprint(os.Stderr, "Invalid input. Please enter something to calculate the sum\n")
		os.Exit(1)
	}

	switch *sha {
	case SHA256:
		fmt.Printf("input: %s\nSHA256: %x\n", *text, sha256.Sum256([]byte(*text)))
	case SHA384:
		fmt.Printf("input: %s\nSHA384: %x\n", *text, sha512.Sum384([]byte(*text)))
	case SHA512:
		fmt.Printf("input: %s\nSHA512: %x\n", *text, sha512.Sum512([]byte(*text)))
	default:
		fmt.Fprintf(os.Stderr, "Invalid SHA %d!\n", *sha)
		os.Exit(1)
	}
}
