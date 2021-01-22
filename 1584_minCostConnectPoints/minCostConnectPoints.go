package leetcode

func minCostConnectPoints(points [][]int) int {

	length := len(points)

	if length == 1 {
		return 0
	} else if length == 2 {
		return cost(points[0], points[1])
	}

	connectCount := 1
	orderIndexs := make([]int, length)
	costs := make([][]int, length)
	totalCost := 0

	for i := 0; i < length; i++ {

		costs[i] = make([]int, length)
		orderIndexs[i] = i

		for j := i - 1; j >= 0; j-- {
			costs[i][j] = costs[j][i]
		}

		for j := i + 1; j < length; j++ {
			costs[i][j] = cost(points[i], points[j])
		}
	}

	for connectCount < length {
		minCost := 2000001
		minIndex := 0

		for j := connectCount; j < length; j++ {
			for i := 0; i < connectCount; i++ {
				if costs[orderIndexs[i]][orderIndexs[j]] < minCost {
					minCost = costs[orderIndexs[i]][orderIndexs[j]]
					minIndex = j
				}
			}
		}

		totalCost += minCost

		if minIndex > connectCount {
			tmp := orderIndexs[connectCount]
			orderIndexs[connectCount] = orderIndexs[minIndex]
			orderIndexs[minIndex] = tmp
		}

		connectCount++
	}

	return totalCost
}

func cost(r []int, l []int) int {

	cost := 0

	if r[0] > l[0] {
		cost += r[0] - l[0]
	} else {
		cost += l[0] - r[0]
	}

	if r[1] > l[1] {
		cost += r[1] - l[1]
	} else {
		cost += l[1] - r[1]
	}

	return cost
}
