package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "\t\tstärt   and tab:\t\tend\nof line"
	fmt.Println(s)

	b := uni_spaces([]byte(s))
	fmt.Println(string(b))
	
}

func uni_spaces(bys []byte) []byte {
	prev_space := false
	read_i := 0
	write_i := 0
	
	for read_i < len(bys) {
		rune, size := utf8.DecodeRune(bys[read_i:])

		if !unicode.IsSpace(rune) {
			utf8.EncodeRune(bys[write_i:], rune)
			write_i += size
			prev_space = false
		} else {
			if !prev_space {
				size = utf8.EncodeRune(bys[write_i:], ' ')
				write_i += size
			}
			prev_space = true
		}

		read_i += size
		
	}

	return bys[:write_i]
}
