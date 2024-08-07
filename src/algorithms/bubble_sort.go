package algorithms

func BubbleSort(bubbleArray []int) []int {
	for i := 0; i < len(bubbleArray); i++ {
		for j := 0; j < len(bubbleArray)-1-i; j++ {
			if bubbleArray[j] > bubbleArray[j+1] {
				tmp := bubbleArray[j]
				bubbleArray[j] = bubbleArray[j+1]
				bubbleArray[j+1] = tmp
			}
		}
	}
	return bubbleArray
}

// NOTE: look into visual examples to help cement the logic
