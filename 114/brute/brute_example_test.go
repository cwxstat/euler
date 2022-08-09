package brute_test

import (
	"fmt"
	"github.com/cwxstat/euler/114/brute"
)

func ExampleB_Valid() {

	b := &brute.B{}
	b.Fill(7)
	for i := 0; i < 15; i++ {
		if b.Valid(i) {
			b.Pr(i)
		}
	}
	// Output:
	// 0000000
	// 0000111
	// 0001110

}

func ExampleB_Fill() {

	for i := 3; i < 10; i++ {
		b := &brute.B{}
		b.Fill(i)
		fmt.Printf("C(%d) = %d\n", i, b.Count())
	}

	// Output:
	// C(3) = 2
	// C(4) = 4
	// C(5) = 7
	// C(6) = 11
	// C(7) = 17
	// C(8) = 27
	// C(9) = 44
}
