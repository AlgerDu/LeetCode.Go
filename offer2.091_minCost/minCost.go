package leetcode

func minCost(costs [][]int) int {

	count := len(costs)

	for i := 1; i < count; i++ {
		for j := 0; j < 3; j++ {
			costs[i][j] = min(costs[i-1][(j+1)%3], costs[i-1][(j+2)%3]) + costs[i][j]
		}
	}

	rst := min(costs[count-1][0], costs[count-1][1])
	rst = min(rst, costs[count-1][2])

	return rst
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
