package algorithms

import "math"

func twoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < (len(breaks)); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	for j := 0; j <= int(jumpAmount) && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}
	return -1
}
