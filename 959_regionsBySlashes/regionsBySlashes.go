package leetcode

func regionsBySlashes(grid []string) int {
	n := len(grid)

	downIndexs := make([]int, n)

	currIndex := 0

	for i := 0; i < n; i++ {

		lastRigthIndex := 0
		unstableJ := -1

		for j := 0; j < n; j++ {

			currD := 0
			currR := 0

			if grid[i][j] == '\\' {
				if lastRigthIndex == 0 {
					currIndex++
					currD = currIndex

				} else {
					currD = lastRigthIndex
				}

				unstableJ = -1

				if downIndexs[j] == 0 {
					currIndex++
					currR = currIndex

				} else {
					currR = downIndexs[j]
				}

			} else if grid[i][j] == '/' {
				if lastRigthIndex == 0 && downIndexs[j] == 0 {
					currIndex++

				} else if lastRigthIndex == 0 || downIndexs[j] == 0 {

				} else if lastRigthIndex == downIndexs[j] {

				} else {
					currIndex--
					for m := j + 1; m < n && downIndexs[m] == downIndexs[j]; m++ {
						downIndexs[m] = lastRigthIndex
					}

					currIndex--
				}

				currIndex++
				count++

				currR = currIndex
				currD = currIndex
			} else {
				if lastRigthIndex == 0 && downIndexs[j] == 0 {
					currIndex++
					count++

					currR = currIndex
					currD = currIndex
				} else if lastRigthIndex == 0 || lastRigthIndex == downIndexs[j] {
					currR = downIndexs[j]
					currD = downIndexs[j]

				} else if downIndexs[j] == 0 {
					currR = lastRigthIndex
					currD = lastRigthIndex
				} else {
					count--
					for m := j + 1; m < n && downIndexs[m] == downIndexs[j]; m++ {
						downIndexs[m] = lastRigthIndex
					}

					currR = lastRigthIndex
					currD = lastRigthIndex
				}
			}

			lastRigthIndex = currR
			downIndexs[j] = currD
		}
	}

	return currIndex
}
