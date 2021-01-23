package leetcode

func makeConnected(n int, connections [][]int) int {

	if n-1 > len(connections) {
		return -1
	}

	groupCount := n
	groupIndex := 0

	pcGroupMarks := make([][]int, n)

	for i := 0; i < len(connections); i++ {
		left := connections[i][0]
		rigth := connections[i][1]

		if pcGroupMarks[left] == nil && pcGroupMarks[rigth] == nil {
			groupIndex++
			groupCount--

			pcGroupMarks[left] = make([]int, 2)
			pcGroupMarks[left][0] = groupIndex
			pcGroupMarks[left][1] = rigth

			pcGroupMarks[rigth] = make([]int, 2)
			pcGroupMarks[rigth][0] = groupIndex
			pcGroupMarks[rigth][1] = -1

			connections[groupIndex-1][0] = left
			connections[groupIndex-1][1] = rigth
		} else if pcGroupMarks[rigth] == nil {
			groupCount--

			leftIndex := pcGroupMarks[left][0]

			endPcIndex := connections[leftIndex-1][1]
			pcGroupMarks[endPcIndex][1] = rigth
			connections[leftIndex-1][1] = rigth

			pcGroupMarks[rigth] = make([]int, 2)
			pcGroupMarks[rigth][0] = leftIndex
			pcGroupMarks[rigth][1] = -1

		} else if pcGroupMarks[left] == nil {
			groupCount--

			rightIndex := pcGroupMarks[rigth][0]

			endPcIndex := connections[rightIndex-1][1]
			pcGroupMarks[endPcIndex][1] = left
			connections[rightIndex-1][1] = left

			pcGroupMarks[left] = make([]int, 2)
			pcGroupMarks[left][0] = rightIndex
			pcGroupMarks[left][1] = -1

		} else if pcGroupMarks[left][0] == pcGroupMarks[rigth][0] {

		} else {
			groupCount--

			leftIndex := pcGroupMarks[left][0]
			rigthIndex := pcGroupMarks[rigth][0]

			pcIndex := connections[rigthIndex-1][0]

			for pcIndex != -1 {
				pcGroupMarks[pcIndex][0] = leftIndex

				pcIndex = pcGroupMarks[pcIndex][1]
			}

			pcGroupMarks[connections[leftIndex-1][1]][1] = connections[rigthIndex-1][0]

			connections[leftIndex-1][1] = connections[rigthIndex-1][1]
		}
	}

	return groupCount - 1
}
