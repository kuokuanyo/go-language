//非負小整數集合
//零值代表空集合
package t

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}
//回報集合是否包含x
//方法宣告
func (s *IntSet) Has(x int) bool {
	//找尋x位元時以x / 64 的商作為字索引
	//以x % 64的餘作為字中的位元索引
	word, bit := x/64, uint(x%64)
	//return bool
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//將x值加入集合
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

//將s設為s與t的聯集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword //聯集的意思 : s.words[i] = s.words[i] | tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//將集合以"{1 2 3}"形式的字串回傳
//方法宣告
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}