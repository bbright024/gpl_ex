package main

import (
	"fmt"
//	"unicode"
	"unicode/utf8"
)

func main() {
	s := "\t\tstärt   and tab:\t\tend\nof line"
	fmt.Println(s)

	b := reverse([]byte(s))
	fmt.Println(string(b))

}
// reverse the characters of a []byte slice that represents a UTF-8 encoded string, in place
func reverse(bys []byte) []byte {
	n := utf8.RuneCount(bys)
	if n <= 1 || !utf8.Valid(bys){
		return bys
	}
	n = len(bys) - 1

	for head_i, tail_i := 0, n + 1; head_i < tail_i; {

		head_rune, head_size := utf8.DecodeRune(bys[head_i:])
		tail_rune, tail_size := utf8.DecodeLastRune(bys[:tail_i])

		if tail_size != head_size { 
			copy(bys[(head_i + tail_size):(tail_i - head_size)],
				bys[(head_i + head_size):(tail_i - tail_size)])
		}
		
		utf8.EncodeRune(bys[head_i:], tail_rune)
		utf8.EncodeRune(bys[tail_i - head_size:], head_rune)

		head_i += tail_size;
		tail_i -= head_size;
	}
	
	return bys
}
