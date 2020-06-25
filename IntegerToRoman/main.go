package IntegerToRoman

import (
	"strings"
)

func Int2Roman(num int) string {
	var ans strings.Builder
	var romanNumeralDict = map[int]string{
		1000: "M",
		900 : "CM",
		500 : "D",
		400 : "CD",
		100 : "C",
		90  : "XC",
		50  : "L",
		40  : "XL",
		10  : "X",
		9   : "IX",
		5   : "V",
		4   : "IV",
		1   : "I",
	}
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9,5,4,1}

	for _, value := range keys {
		if num < value {
			continue
		}
		var count int = num / value
		if count > 0 {
			for i:=0; i< count; i++ {
				ans.WriteString(romanNumeralDict[value])
			}
			num = num - count * value
		}
	}
	return ans.String()
}