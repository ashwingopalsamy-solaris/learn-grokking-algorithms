package main

import (
	"fmt"
	"math/rand"
	"time"

	binary_search "github.com/ashwin2125/learn-grokking-algorithms/exercises/chapter-01/binary-search/maximum-steps-binary-search-n-input/pkg/binary-search"
)

func main() {
	fmt.Println("Binary Search Implementation")
	rand.Seed(time.Now().UnixNano())

	// Initialize an empty slice to hold the random values
	slice128 := make([]int, 128)
	for i := 0; i < len(slice128); i++ {
		slice128[i] = i + 1 // Generates random integers between 0 and 100
	}

	fmt.Println("Number of steps taken is: ", binary_search.Binary_search(slice128, 128))

	slice256 := make([]int, 256)
	for i := 0; i < 256; i++ {
		slice256[i] = i + 1
	}

	fmt.Println("Number of steps taken is: ", binary_search.Binary_search(slice256, 256))
}
