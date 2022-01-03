package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

const UINT_SIZE = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	for i := range s.words {
		copy.words = append(copy.words, s.words[i])
	}
	return &copy
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, 0)
		}
	}
	sLength := len(s.words)
	tLength := len(t.words)
	if sLength > tLength {
		for i := tLength; i < sLength; i++ {
			s.words[i] = 0
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmeticDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}

func (s *IntSet) AddAll(ints ...int) {
	for _, val := range ints {
		s.Add(val)
	}
}

//string in form {1, 3, 4}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				//put a space before every number except the first
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UINT_SIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for i := 0; i < UINT_SIZE; i++ {
			if word&1 != 0 {
				count++
			}
			word = word >> 1
		}
	}
	return count
}

func (s *IntSet) Elems() (result []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, UINT_SIZE*i+j)
			}
		}
	}
	return result
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(123)
	x.Add(5)
	fmt.Println(x.String(), " size=", x.Len())
	y.Add(2)
	y.Add(5)
	fmt.Println(y.String(), " size=", y.Len())
	x.UnionWith(&y)
	fmt.Println(x.String(), " size=", x.Len())
	x.Remove(123)
	fmt.Println(x.String())
	x.Clear()
	fmt.Println(x.String())
	x = *y.Copy()
	fmt.Println(x.String())
	x.AddAll(8, 12)
	fmt.Println(x.String())
	x.IntersectWith(&y)
	fmt.Println(x.String())
	y.Remove(5)
	fmt.Printf("%s - %s = ", x.String(), y.String())
	x.DifferenceWith(&y)
	fmt.Println(x.String())
	x.AddAll(5, 7)
	y.AddAll(7, 9)
	fmt.Printf("%s x %s = ", x.String(), y.String())
	x.SymmeticDifferenceWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Elems())
}
