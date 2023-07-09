package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	ex := flag.Int("ex", 1, "Number of the functionality to run 1 or 2")
	flag.Parse()

	switch *ex {
	case 1:
		charCount()
	case 2:
		wordFreq()
	default:
		fmt.Fprintln(os.Stderr, "Invalid functionality!")
		os.Exit(1)
	}
}

type runeTest func(rune) bool

var runeTests = map[string]runeTest{
	"Digit":   unicode.IsDigit,
	"Letter":  unicode.IsLetter,
	"Lower":   unicode.IsLower,
	"Upper":   unicode.IsUpper,
	"Control": unicode.IsControl,
	"Graphic": unicode.IsGraphic,
	"Mark":    unicode.IsMark,
	"Number":  unicode.IsNumber,
	"Print":   unicode.IsPrint,
	"Punct":   unicode.IsPunct,
	"Space":   unicode.IsSpace,
	"Symbol":  unicode.IsSymbol,
	"Title":   unicode.IsTitle,
}

func charCount() {
	fmt.Println("\ncharcount()")

	categories := make(map[string]int)

	for k := range runeTests {
		categories[k] = 0
	}

	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		testRune(categories, r)
	}

	fmt.Printf("\n%-15s\t%10s\n", "Categories", "Count")

	for cat, count := range categories {
		fmt.Printf("%-15s\t%10d\n", cat, count)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func testRune(categories map[string]int, r rune) {
	for k, f := range runeTests {
		categories[k] += boolToInt(f(r))
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func wordFreq() {

	fmt.Println("\nwordFreq()")

	freq := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		freq[word]++
	}

	fmt.Printf("\n%-20s\t%10s\n", "Words", "Count")

	for word, count := range freq {
		fmt.Printf("%-20s\t%10d\n", word, count)
	}
}
