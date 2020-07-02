package TwoCityScheduling

import (
	"awesomeProject/Common"
	"fmt"
	"sort"
)

func TwoCityCost(costs [][]int) int {
	sort.Slice(costs, func (i,j int) bool {
		return Common.Abs(costs[i][0] - costs[i][1]) > Common.Abs(costs[j][0] - costs[j][1])
	})

	var ans = 0
	n := len(costs) / 2
	i,j := 0, 0
	for _, cost := range costs {
		if i >= n {
			ans += cost[0]
		}
		if j >= n {
			ans += cost[1]
		}
		if i < n && j < n {
			if cost[0] > cost[1] {
				i++
				ans += cost[1]
			} else {
				j++
				ans += cost[0]
			}
		}
	}
	fmt.Printf("%#v", costs)

	return ans
}