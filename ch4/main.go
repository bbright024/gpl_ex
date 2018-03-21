// my solutions, chapter 4 exercises

package main

import (
	"fmt"
)

const SIX = 6

func main() {
	a := [SIX]int{0, 1, 2, 3, 4, 5}
	rev2(&a)
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(2, s)
	fmt.Println(s)

	strs := []string{"hello", "world", "world", "world", "a", "b", "a", "b", "b"}
	fmt.Println(strs)
	strs = remdup(strs)
	fmt.Println(strs)

	strs = []string{"hello", "world", "world", "world", "a", "b", "a", "b", "b"}
	fmt.Println(strs)
	strs = remdup2(strs)
	fmt.Println(strs)
}

//4.3
func rev2(a *[SIX]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

//4.4
//rotate s by n
func rotate(n int, s []int) []int {
	temp := []int{}
	temp = append(temp, s[n:]...)
	temp = append(temp, s[:n]...)

	return temp

}

//4.5a
// not in place, allocs new slice
func remdup(strs []string) []string {
	out := []string{strs[0]}
	i := 0
	for _, s := range strs[1:] {
		if out[i] != s {
			out = append(out, s)
			i++
		}

		fmt.Println(out)
	}
	return out
}

//4.5b
//in-place dup removal
func remdup2(strs []string) []string {
	n := len(strs)
	if n <= 1 {
		return strs
	}

	uniq := 1
	prev := strs[0]

	for i := 1; i < n; i++ {
		next := strs[i]
		if prev != next {
			uniq++
			prev = next
		} else {
			copy(strs[i:], strs[i+1:])
			n--
			i--
		}
		fmt.Println(prev)
	}

	return strs[:uniq]

}
