package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var testSlice []int // Empty slice for comparison
	var testSlice2 = make([]int, 0, n) // Preallocated slice with capacity n

	// Time the append operation without preallocation
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))

	// Reinitialize testSlice to ensure it's empty before the second test
	testSlice = []int{} 

	// Time the append operation with preallocation
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()

	// Append elements until the slice has n elements
	for len(slice) < n {
		slice = append(slice, 1)
	}

	// Return the elapsed time
	return time.Since(t0)
}
