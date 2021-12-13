package main

import (
	"fmt"
	"time"
)

// var LongestChain = make(map[int]int)

type LongestChainResult struct {
	Number int
	Chain  int
}

// Collatz returns the number of steps required to reach 1 from the given number
func Collatz(n int, result *LongestChainResult) int {
	chain := 0
	calculate := n
	for n != 1 {
		n = []int{n >> 1, 3*n + 1}[n&1]
		chain++
	}

	if chain > result.Chain {
		result.Chain = chain
		result.Number = calculate
	}

	return chain
}

func main() {
	fmt.Printf("Start: %s \n", time.Now().Format("2006-01-02 15:04:05.000"))

	result := new(LongestChainResult)
	for i := 1; i < 100000000; i++ {
		_ = Collatz(i, result)
	}

	fmt.Printf("The longest chain is %d and the number is %d \n", result.Chain, result.Number)
	fmt.Printf("End:   %s \n", time.Now().Format("2006-01-02 15:04:05.000"))
}
