package algorithms

func LinearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

// next algorithm is binary search
