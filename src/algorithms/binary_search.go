package algorithms

func BinarySearch(array []int, lowValue int, highValue int, target int) bool {
	midPoint := (lowValue + (highValue-lowValue)/2)
	value := array[midPoint]

	if lowValue > highValue {
		return false
	}
	if value == target {
		return true
	} else if value > target {
		return BinarySearch(array, lowValue, midPoint-1, target)
	} else {
		return BinarySearch(array, midPoint+1, highValue, target)
	}
}

// TODO: practice other binary search related problems, you need practice with this because
// you had to look up guides to do it, plus you got it wrong

// TODO: implement it without doing recursion now.
