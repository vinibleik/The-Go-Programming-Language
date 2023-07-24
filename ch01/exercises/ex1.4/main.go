// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		if err := countLines(os.Stdin, counts, foundIn); err != nil {
			log.Fatal("ex1.4: ", err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			if err := wrapOpenFile(arg, counts, foundIn); err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
				continue
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Count: %4d\tFound In: %v\tLine: %s\n", n, foundIn[line], line)
		}
	}
}

func wrapOpenFile(file string, counts map[string]int, foundIn map[string][]string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := countLines(f, counts, foundIn); err != nil {
		return err
	}
	return nil
}

func in(curFile string, files []string) bool {
	for _, file := range files {
		if curFile == file {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(f.Name(), foundIn[line]) {
			foundIn[line] = append(foundIn[line], f.Name())
		}
	}

	return input.Err()
}
