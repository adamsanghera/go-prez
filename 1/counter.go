package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	numCounters = int64(4)
)

// counter calculates the gaussian sum range (start, end].
// Therefore, usage should be (0, 10], (10, 20), etc.
func counter(start, end, id int64, outbox chan int64) {
	result := int64(0)
	for idx := start; idx < end; idx++ {
		result += idx
		// log.Printf("%d counts %d\n", id, idx)
	}
	outbox <- result
}

func main() {
	startTime := time.Now()
	defer log.Printf("Finished execution in %d ns", time.Since(startTime).Nanoseconds())

	// Get N from the passed-in arg
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Make our return channel
	inbox := make(chan int64)
	numAliveRoutines := 0

	// Spawn our counters
	for idx := int64(0); idx < numCounters; idx++ {
		start := (idx * int64(n) / numCounters) + 1
		end := ((idx + 1) * int64(n) / numCounters) + 1

		// In the case where the number is less than 10
		// we can avoid spawning 10 routines = nice.
		if start < end {
			numAliveRoutines++
			go counter(start, end, idx, inbox)
		}
	}

	// Collect our results
	finalResult := int64(0)
	for id := 0; id < numAliveRoutines; id++ {
		log.Printf("Blocking on reply from a worker...")
		piece := <-inbox
		log.Printf("Received partial sum %d\n", piece)
		finalResult += piece
	}

	// Print the result
	fmt.Printf("Final final result: %d\n", finalResult)
}
