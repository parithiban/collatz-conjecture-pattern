package main

import (
	"fmt"
	"time"
)

// Collatz returns the number of steps required to reach 1 from the given number
func Collatz(n int) int {
	chain := 0
	for n != 1 {
		n = []int{n >> 1, 3*n + 1}[n&1]
		chain++
	}

	return chain
}

func main() {
	fmt.Printf("Start: %s \n", time.Now().Format("2006-01-02 15:04:05.000"))

	for i := 1; i < 100000000; i++ {
		//Skipping the power of 2 numbers as it will awlays end in 8,4,2,1
		// if (i & (i - 1)) == 0 {
		// 	continue
		// }

		_ = Collatz(i)
	}

	fmt.Printf("End:   %s \n", time.Now().Format("2006-01-02 15:04:05.000"))
}
