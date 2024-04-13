package main

import (
	"fmt"

	binary_search "github.com/ashwin2125/learn-grokking-algorithms/exercises/chapter-01/binary-search/binary-search-implementation-go/pkg/binary-search"
)

func main() {
	fmt.Println("Binary Search Implementation")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Index of the Target Item is: ", binary_search.Binary_search(slice, 5))
}
