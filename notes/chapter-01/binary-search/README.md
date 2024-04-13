# Binary Search

## Notes:

- One of the most widely used searching algorithm.
- Only works if the given list / array is sorted.
- In general, for any list of `n`, the number of steps taken would be `log2n` in the worst case, whereas the simple search takes `n` steps.

## Syntax:

```go
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
```