package leetcode

func isRectangleCover(rectangles [][]int) bool {

	rectangleCount := len(rectangles)
	lineCount := rectangleCount * 4
	edges := make([]bool, 4)
	lineChecked := make([]bool, lineCount)

	paraleLines := make([]int, rectangleCount)
	paraleLines2 := make([]int, rectangleCount)

	for i := 0; i < lineCount; i++ {

		if lineChecked[i] {
			continue
		}

		paraleLines[0] = i

		edgeIndex := i % 4
		rectangleIndex := i / 4
		count := 1
		count2 := 0

		m := (4 - edgeIndex) % 4
		n := rectangles[rectangleIndex][m]

		for j := rectangleIndex + 1; j < rectangleCount; j++ {

			checkLineIndex := rectangleIndex*4 + edgeIndex

			if !lineChecked[checkLineIndex] && rectangles[j][m] == n {
				lineChecked[checkLineIndex] = true
				count++
				paraleLines[count] = checkLineIndex
				continue
			}

			checkLineIndex = rectangleIndex*4 + (edgeIndex+2)%4

			if !lineChecked[checkLineIndex] && rectangles[j][(m+2)%4] == n {
				lineChecked[checkLineIndex] = true
				count2++
				paraleLines2[count2] = checkLineIndex
			}
		}

		if count2 == 0 && edges[edgeIndex] {
			// 找到了重复的边
			return false
		}
	}

	return true
}
