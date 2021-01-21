package leetcode

func minCostConnectPoints(points [][]int) int {

	length := len(points)
	costs := make([][]int, length)
	groupMarks := make([]int, length)
	groupCosts := make([]int, length/2+1)
	groupStartPoint := make([]int, length/2+1)
	groupIndex := 1

	for i := 0; i < length; i++ {

		costs[i] = make([]int, length)

		minCost := 2000000
		minJ := 0

		for j := i - 1; j >= 0; j-- {
			costs[i][j] = costs[j][i]

			if costs[i][j] < minCost {
				minCost = costs[i][j]
				minJ = j
			}
		}

		for j := i + 1; j < length; j++ {
			cost := cost(points[i], points[j])

			if cost < minCost {
				minCost = cost
				minJ = j
			}

			costs[i][j] = cost
		}

		if groupMarks[i] == 0 && groupMarks[minJ] == 0 {
			groupMarks[i] = groupIndex
			groupMarks[minJ] = groupIndex

			costs[i][i] = minJ

			groupCosts[groupIndex] = minCost
			groupStartPoint[groupIndex] = i

			groupIndex++
		} else if groupMarks[i] == 0 {
			groupMarks[i] = groupMarks[minJ]
			costs[minJ][minJ] = i

			groupCosts[groupMarks[minJ]] += minCost
		} else if groupMarks[minJ] == 0 {
			groupMarks[minJ] = groupMarks[i]
			costs[i][i] = minJ

			groupCosts[groupMarks[i]] += minCost
		} else {
			minIndex := groupMarks[i]
			maxIndex := groupMarks[i]

			if minIndex > groupMarks[minJ] {
				minIndex = groupMarks[minJ]
			} else {
				maxIndex = groupMarks[minJ]
			}

			groupCosts[minIndex] += minCost + groupCosts[maxIndex]

			end := i

			for costs[end][end] != 0 {
				end = costs[end][end]
			}

			costs[end][end] = groupStartPoint[maxIndex]

			for costs[end][end] != 0 {
				end = costs[end][end]

				groupMarks[end] = minIndex
			}

			groupIndex--
		}
	}

	return groupCosts[1]
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
