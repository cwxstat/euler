package brute
import (
	"fmt"
	"math"
)

type B struct {
	s   []string
	len int
}

func (b *B) Pr(j int) *B {
	fmt.Printf("%s\n", b.s[j])
	return b
}

func (b *B) Count() int {
	count := 0
	for i := 0; i < len(b.s); i++ {
		if b.Valid(i) {
			count++
			//b.Pr(i)
		}
	}
	return count
}

func (b *B) Valid(j int) bool {
	var count int
	for i := 0; i < b.len; i++ {
		if b.s[j][i] == '1' {
			count++
		} else {
			if count == 1 || count == 2 {
				return false
			}
			count = 0
		}

	}
	if count == 1 || count == 2 {
		return false
	}
	return true
}

func (b *B) Fill(j int) *B {
	b.len = j
	myfmt := fmt.Sprintf("%%0%db", j)

	for i := uint64(0); i < uint64(math.Pow(2, float64(j))); i++ {
		b.s = append(b.s, fmt.Sprintf(myfmt, i))
	}
	return b
}

