// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
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
		/*
// I'm assuming the authors didn't mean to do EVERY Is...
//  but it was nice to read about them all
		if unicode.IsControl(c) {
				}
		if unicode.IsDigit(c) {
		
		}
		if unicode.IsGraphic(c) {
		
		}
		if unicode.IsLetter(c) {
		
		}
		if unicode.IsLower(c) {
		
		}
		if unicode.IsMark(c) {
		
		}
		if unicode.IsNumber(c) {
		
		}
		if unicode.IsPrint(c) {
		
		}
		if !unicode.IsPrint(c) {
		
		}
		if unicode.IsPunct(c) {
		
		}
		if unicode.IsSpace(c) {
		
		}
		if unicode.IsSymbol(c) {
		
		}
		if unicode.IsTitle(c) {
		
		}
		if unicode.IsUpper(c) {
		
		}
*/		
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
