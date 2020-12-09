package method

import (
	"bytes"
	"fmt"
)

func BitsetEntry() {
	s := IntSet{}
	s.Add(1)
	s.Add(7)
	s.Add(14)
	fmt.Printf("s = %v\n", s.String())

	t := IntSet{}
	t.Add(100)
	t.Add(200)
	t.Add(9)
	fmt.Printf("t = %v\n", t.String())

	s.UnionWith(&t)
	fmt.Printf("s + t = %v\n", s)
	fmt.Printf("s + t = %v\n", &s)
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
