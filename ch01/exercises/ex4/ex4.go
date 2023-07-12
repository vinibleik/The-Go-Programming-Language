package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupEx4: %v\n", err)
				continue
			}

			countLines(file, counts)
			file.Close()
		}
	}
	printLines(counts)
}

func countLines(file *os.File, counts map[string]map[string]int) {
	buffer := bufio.NewScanner(file)
	file_map := make(map[string]int)
	for buffer.Scan() {
		file_map[buffer.Text()]++
	}
	counts[file.Name()] = file_map
}

func printLines(counts map[string]map[string]int) {
	for file_name, file_map := range counts {
		for line, count := range file_map {
			if count > 1 {
				fmt.Printf("File: %s Line: %s Count: %d%c", file_name, line, count, 10)
			}
		}
	}
}
