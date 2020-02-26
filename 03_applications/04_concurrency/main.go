package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// We start a concurrent count over steps number of iterations with cnt as increments
	for steps, increment := range map[int]int {
		100: 5,
		//10000: 1,
		//100_000: 1,
	}{
		cnt := countConcurrent(steps, increment)
		logResult(steps, increment, cnt)
	}
}

func logResult(steps int, increment int, cnt int) {
	fmt.Println("Calling countCuncurrent for %d steps and increment %d", steps, increment)
	fmt.Println("------------------------------------------------------")
	fmt.Println("Value returned from function back to main as: ", cnt)
}

func countConcurrent (steps int, increment int) int {
	input := make(chan int) // Collects results from countUp for aggregation
	results := make(chan int) // Channel to collect final result after aggregation

	var wg sync.WaitGroup
	log.Println("Starting handleCountResults")
	go  handleCountResult(&wg, input, results)

	//count for number of steps concurrently
	for i := 1; i <= steps; i++ {
		log.Printf("Sending off count %d", i)
		wg.Add(1)
		go countUp(increment, input)
	}

	// PAUSE: we need to wait for all gophers to return their jobs before continuing
	log.Printf("Waiting for all go routines to return")
	wg.Wait()
	// Now they have all returned their responses. We can shut down the input and send the close
	// signal to the handleCount to release the range loop
	log.Printf("all have returned")
	close(input)

	// Now we can collect the final result from the aggregation function (handleCountResult)
	res := <- results
	log.Printf("result received as %v", res)
	return res
}

// We launch steps number of individual counters. Each counter does something and returns an increment of n
func countUp(n int, out chan int) {
	// This function will send a number to the channel
	time.Sleep(time.Duration(1*time.Millisecond)) // Do some work
	out <- n
	return
}

// Handlecount idles while waiting to receive numbers from the input channel send by the countUp functions.
// It uses the WaitGroup to know how many results to wait for before returning the aggregate (After seeing all counts)
func handleCountResult(wg *sync.WaitGroup, input chan int, output chan int) {
	cnt := 0
	// Range over input channel to receive counts
	for n := range input {
		log.Printf("GO: increating cnt:%d by %d", cnt, n)
		cnt += n
		// counting down received messages
		wg.Done()
	}
	// After all are done, the loop returns with a break and we can safely return the result
	// to the provided output channel
	output <- cnt
}

