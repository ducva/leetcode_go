package containermostwater

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxArea(height []int) int {
	// return the container with most water
	ans := 0
	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j > i; j-- {
			temp := min(height[i], height[j]) * (j - i)
			if temp > ans {
				ans = temp
			}
		}
	}
	return ans
}
