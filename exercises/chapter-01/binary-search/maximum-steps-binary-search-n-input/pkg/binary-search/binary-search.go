package handler

func Binary_search(slice []int, target int) int {
	if len(slice) == 0 {
		return -1
	}

	low, high := 0, len(slice)-1
	stepsTaken := 0

	for low <= high {
		mid := low + (high-low)/2
		guess := slice[mid]
		if guess == target {
			return stepsTaken
		} else if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
		stepsTaken = stepsTaken + 1
	}
	return stepsTaken
}
