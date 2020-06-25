package DivideTwoIntegers

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Divide(dividend, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var ans = 0
	signed := (divisor > 0 && dividend < 0) || (dividend > 0 && divisor < 0)
	dividendAbs := abs(dividend)
	div := abs(divisor)
	temp := 0
	for i:=31; i >= 0; i-- {
		if temp + div << i <= dividendAbs {
			temp += div << i
			ans |= 1 << i
		}
	}

	if signed {
		return max(0 - ans, math.MinInt32 - 1)
	} else {
		return min(ans, math.MaxInt32)
	}
}

func Add(x, y int) int {
	if y == 0 {
		return x
	}
	return Add(x ^ y, x & y << 1)
}

func Subtract(x, y int) int {
	if y == 0 {
		return x
	}
	return Subtract(x ^ y, (^x) & y << 1)
}

func Mul(x, y int) int {
	/*
	Do multiplication by implementing Russian Peasant algorithms.
	- Represent multiplicand and multiplier in binary format, then we can see
	the fact is we multiple the multiplicand with each bit of multiplier.
	In case the bit of multiplier is 0, the result will be ignored
	 */
	if y == 0  || x == 0 {
		return 0
	}
	ans := 0
	signed := (x > 0 && y < 0) || (x < 0 && y > 0)
	multiplier, multiplicand := abs(y), abs(x)
	for {
		if multiplier & 1 > 0 {
			ans += multiplicand
		}
		multiplier = multiplier >> 1
		multiplicand = multiplicand << 1
		if multiplier <= 0 {
			break
		}
	}
	if signed {
		return 0 - ans
	}
	return ans
}