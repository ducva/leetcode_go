package ArrangingCoins

import "math"

func ArrangingCoins(n int) int {
	ans := 0
	if n > 0 {
		ans = int(math.Floor(math.Sqrt(float64((float64(2*n))+0.25)) - 0.5))
	}
	return ans
}
