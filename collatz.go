package collatz

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Producers is a channel of producers
type Producer struct {
	Start int
	End   int
}

//Consumer is a channel of consumers
type Consumer struct {
	Chain  int
	Number int
}

var Producers = make(chan Producer, 50)
var Consumers = make(chan Consumer, 50)

// Collatz returns the number of steps required to reach 1 from the given number
func Collatz(n int, result *Consumer) *Consumer {
	chain := 0
	calculate := n
	// internal := make(map[int]int)

	for n != 4 && n >= calculate {
		n = []int{n >> 1, 3*n + 1}[n&1]
		chain++
	}

	if chain > result.Chain {
		result.Chain = chain
		result.Number = calculate
	}

	return result
}

// ConsumeProcess consumes the results
func ConsumeProcess(result chan *Consumer, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	findInConsumers := new(Consumer)
	for consume := range Consumers {
		if consume.Chain > findInConsumers.Chain {
			findInConsumers.Chain = consume.Chain
			findInConsumers.Number = consume.Number
		}
	}

	result <- &Consumer{findInConsumers.Chain, findInConsumers.Number}
}

// Execute executes the collatz algorithm
func Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	for produce := range Producers {
		consumer := new(Consumer)
		for i := produce.Start; i <= produce.End; i++ {
			_ = Collatz(i, consumer)
		}

		Consumers <- Consumer{consumer.Chain, consumer.Number}
	}
}

// InitWorkers initializes the workers
func InitWorkers(workers int) {
	var wg sync.WaitGroup
	defer close(Consumers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go Execute(&wg)
	}
	wg.Wait()
}

// SetUpProducers sets up the producers
func SetUpProducers(input int) {
	defer close(Producers)
	dividend, remainder := input/100, input%100
	i := 1
	for i <= input {
		start := i
		if remainder != 0 && (i-1 == dividend*100) {
			Producers <- Producer{start, i - 1 + remainder}
			break
		}

		i = i + dividend
		Producers <- Producer{start, i - 1}
	}
}

// main is the entry point of the program
func main() {
	var input int
	fmt.Print("Enter the number: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		panic(err)
	}

	start := time.Now()
	go SetUpProducers(input)

	var m sync.Mutex
	result := make(chan *Consumer)

	for i := 0; i < 10; i++ {
		go ConsumeProcess(result, &m)
	}

	workers := runtime.NumCPU()
	InitWorkers(workers)

	longestChain := <-result

	fmt.Printf("The longest chain is %d and the number is %d \n", longestChain.Chain, longestChain.Number)
	fmt.Printf("Total Execution time  %d ms \n", time.Since(start).Milliseconds())
}
