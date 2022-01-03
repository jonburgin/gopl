package main

import (
	"fmt"
	"unicode"
)
import "unicode/utf8"

func reverse(arr *[6]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func rotate(s []int, i int) []int {
	left := s[:i]
	right := s[i:]
	for _, c := range left {
		right = append(right, c)
	}
	return right
}

func dedup(strings []string) []string {
	i := 0
	for index, s := range strings {
		if index == 0 {
			continue
		}
		if s != strings[index-1] {
			i++
			strings[i] = s
		}
	}
	strings[i] = strings[len(strings)-1]
	i++
	return strings[:i]
}

func squash(s string) string {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if unicode.IsSpace(r) {
			s = s[:i] + s[i+size:]
		} else {
			i += size
		}
	}
	return s
}

func copy(s []byte, to int, from int, size int) {
	for i := 0; i < size; i++ {
		s[to] = s[from]
		to++
		from++
	}
}

func squashByte(s []byte) []byte {
	removed := 0
	for to, from := 0, 0; from < len(s); {
		r, size := utf8.DecodeRune(s[from:])
		if unicode.IsSpace(r) {
			removed += size
		} else {
			copy(s, to, from, size)
			to += size
		}
		from += size
	}
	return s[:len(s)-removed]
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	reverse(&arr)
	fmt.Println(arr)
	fmt.Println(rotate([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(dedup([]string{"How", "now", "now", "brown", "cow", "cow"}))
	fmt.Println(squash("I don't need    no spaces"))
	output := squashByte([]byte("I don't need    no spaces"))
	fmt.Println(string(output))
}
