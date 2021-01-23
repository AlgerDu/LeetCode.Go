package leetcode

func makeConnected(n int, connections [][]int) int {

	if n-1 > len(connections) {
		return -1
	}

	connectionMarks := make([]int, n)
	connectionCount := 0

	for i := 0; i < len(connections); i++ {
		for j := 0; j < 2; j++ {
			if connectionMarks[connections[i][j]] != 1 {
				connectionCount++
				connectionMarks[connections[i][j]] = 1
			}
		}
	}

	return n - connectionCount
}
