package leetcode

func minCostConnectPoints(points [][]int) int {

	length := len(points)

	if length == 1 {
		return 0
	} else if length == 2 {
		return cost(points[0], points[1])
	}

	costs := make([][]int, length)
	groupMarks := make([]int, length)
	groupChain := make([]int, length)
	groupCosts := make([]int, length/2+1)
	groupStartPoint := make([]int, length/2+1)
	groupIndex := 0

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

		if groupMarks[i] == groupMarks[minJ] && groupMarks[i] != 0 {

		} else if groupMarks[i] == 0 && groupMarks[minJ] == 0 {
			groupIndex++

			groupMarks[i] = groupIndex
			groupMarks[minJ] = groupIndex

			groupChain[i] = minJ
			groupChain[minJ] = -1

			groupCosts[groupIndex] = minCost
			groupStartPoint[groupIndex] = i

		} else if groupMarks[i] == 0 {
			groupMarks[i] = groupMarks[minJ]

			end := minJ
			for groupChain[end] != -1 {
				end = groupChain[end]
			}
			groupChain[end] = i

			groupChain[i] = -1

			groupCosts[groupMarks[minJ]] += minCost
		} else if groupMarks[minJ] == 0 {
			groupMarks[minJ] = groupMarks[i]

			end := i
			for groupChain[end] != -1 {
				end = groupChain[end]
			}
			groupChain[end] = minJ

			groupChain[minJ] = -1

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

			for groupChain[end] != -1 {
				end = groupChain[end]
			}

			groupChain[end] = groupStartPoint[maxIndex]

			for groupChain[end] != -1 {
				end = groupChain[end]

				groupMarks[end] = minIndex
			}

			for maxIndex < groupIndex {
				groupStartPoint[maxIndex] = groupStartPoint[maxIndex+1]

				maxIndex++
			}

			groupIndex--
		}
	}

	for groupIndex > 1 {
		minCost := 2000000
		minIndex := groupIndex - 1

		for compareIndex := groupIndex - 1; compareIndex > 0; compareIndex-- {
			i := groupStartPoint[groupIndex]

			for i != -1 {
				j := groupStartPoint[compareIndex]

				for j != -1 {
					if costs[i][j] < minCost {
						minCost = costs[i][j]
						minIndex = compareIndex
					}

					j = groupChain[j]
				}

				i = groupChain[i]
			}
		}

		groupCosts[minIndex] += minCost + groupCosts[groupIndex]

		end := groupStartPoint[minIndex]

		for groupChain[end] != -1 {
			end = groupChain[end]
		}

		groupChain[end] = groupStartPoint[groupIndex]

		groupIndex--
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
