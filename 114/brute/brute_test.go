package brute

import (
	"fmt"
	"testing"
)

func TestB_Fill(t *testing.T) {

	for i := 3; i < 10; i++ {
		b := &B{}
		b.Fill(i)
		fmt.Printf("C(%d) = %d\n", i, b.Count())
	}

}
