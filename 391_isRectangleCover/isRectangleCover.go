package leetcode

func isRectangleCover(rectangles [][]int) bool {

	rectangleCout := len(rectangles)

	if rectangleCout == 1 {
		return true
	}

	scanLineYs := make([]int, 2)
	scanLineX := 0

	for i := 0; i < rectangleCout; i++ {
		minIndex := i

		for j := i + 1; j < rectangleCout; j++ {

			if rectangles[j][0] == rectangles[minIndex][0] &&
				rectangles[j][1] == rectangles[minIndex][1] {
				return false
			}

			if rectangles[j][0] < rectangles[minIndex][0] {
				minIndex = j
			}
		}

		if minIndex != i {
			tmp := rectangles[i]
			rectangles[i] = rectangles[minIndex]
			rectangles[minIndex] = tmp
		}
	}

	scanLineX = rectangles[0][0]
	scanLineYs[0] = rectangles[0][1]
	scanLineYs[1] = rectangles[0][3]

	for i := 1; i < rectangleCout; i++ {
		if rectangles[i][3] > scanLineYs[1] {
			scanLineYs[1] = rectangles[i][3]
		}

		if rectangles[i][1] < scanLineYs[0] {
			scanLineYs[0] = rectangles[i][1]
		}
	}

	for i := 0; i < rectangleCout; {

		for j := i; j < rectangleCout && rectangles[j][0] == scanLineX; j++ {

			minIndex := j

			for k := j + 1; k < rectangleCout && rectangles[k][0] == scanLineX; k++ {
				if rectangles[k][1] < rectangles[minIndex][1] {
					minIndex = k
				}
			}

			if minIndex != j {
				tmp := rectangles[j]
				rectangles[j] = rectangles[minIndex]
				rectangles[minIndex] = tmp
			}
		}

		nextScanLineX := rectangles[i][2]
		y := scanLineYs[0]

		j := i
		for ; j < rectangleCout && rectangles[j][0] == scanLineX; j++ {
			if rectangles[j][1] == y {
				y = rectangles[j][3]
			} else {
				return false
			}

			if rectangles[j][2] < nextScanLineX {
				nextScanLineX = rectangles[j][2]
			}
		}

		if y != scanLineYs[1] {
			return false
		}

		scanLineX = nextScanLineX

		for k := i; k < j; k++ {
			if rectangles[k][2] == scanLineX {

				if k != i {
					rectangles[k] = rectangles[i]
				}
				rectangles[i] = nil
				i++

			} else {
				rectangles[k][0] = scanLineX
			}
		}
	}

	return true
}
