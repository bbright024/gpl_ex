package main

import (
	"fmt"
//	"io"
	"os"
	"bufio"

)


func main() {
	wordfreq := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		wordfreq[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word:\tfreq:\n")
	for key, val := range wordfreq {
		fmt.Printf("%q\t%d\n", key, val)
	}
}
