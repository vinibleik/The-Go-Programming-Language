// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type CountLines struct {
	count   int
	foundIn []string
}

func main() {

	counts := make(map[string]*CountLines)

	files := os.Args[1:]
	if len(files) == 0 {
		if err := countLines(os.Stdin, counts); err != nil {
			log.Fatal("ex1.4: ", err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			if err := wrapOpenFile(arg, counts); err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
				continue
			}
		}
	}

	for line, data := range counts {
		if data.count > 1 {
			fmt.Printf("Count: %4d\tFound In: %v\tLine: %s\n", data.count, data.foundIn, line)
		}
	}
}

func wrapOpenFile(file string, counts map[string]*CountLines) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := countLines(f, counts); err != nil {
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

func countLines(f *os.File, counts map[string]*CountLines) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		_, ok := counts[line]
		if !ok {
			counts[line] = new(CountLines)
		}
		counts[line].count++
		if !in(f.Name(), counts[line].foundIn) {
			counts[line].foundIn = append(counts[line].foundIn, f.Name())
		}
	}

	return input.Err()
}
