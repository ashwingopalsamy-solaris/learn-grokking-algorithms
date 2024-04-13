package binary_search

import (
	"fmt"
)

func Binary_search[T comparable](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	low, high := 0, len(slice)-1

	for low <= high {
		mid := low + (high-low)/2
		guess := slice[mid]
		if guess == target {
			return mid
		} else if fmt.Sprintf("%v", guess) < fmt.Sprintf("%v", target) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
