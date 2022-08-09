package brute_test

import (
	"fmt"
	"github.com/cwxstat/euler/114/brute"
)

func ExampleB_Fill()  {

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
