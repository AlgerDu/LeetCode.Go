package leetcode

func isRectangleCover(rectangles [][]int) bool {

	rectangleCount := len(rectangles)
	edgeCount := rectangleCount * 4

	resultEdges := make([]bool, 4)
	edgeCheckState := make([]bool, edgeCount)

	group1 := make([]int, rectangleCount)
	group2 := make([]int, rectangleCount)

	for i := 0; i < edgeCount; i++ {

		if edgeCheckState[i] {
			continue
		}

		group1[0] = i

		edgeCode := i % 4
		rectangleIndex := i / 4

		group1Count := 1
		group2Count := 0

		vIndex := (4 - edgeCode) % 4
		v := rectangles[rectangleIndex][vIndex]

		for j := rectangleIndex + 1; j < rectangleCount; j++ {

			toCheckEdgeIndex := rectangleIndex*4 + edgeCode

			if !edgeCheckState[toCheckEdgeIndex] && rectangles[j][vIndex] == v {
				edgeCheckState[toCheckEdgeIndex] = true
				group1[group1Count] = toCheckEdgeIndex
				group1Count++
				continue
			}

			toCheckEdgeIndex = rectangleIndex*4 + (edgeCode+2)%4

			if !edgeCheckState[toCheckEdgeIndex] && rectangles[j][(vIndex+2)%4] == v {
				edgeCheckState[toCheckEdgeIndex] = true
				group2[group1Count] = toCheckEdgeIndex
				group2Count++
			}
		}

		if group2Count == 0 && resultEdges[edgeCode] {
			return false
		}

		vStart := (edgeCode + 1) / 4
		vEnd := (edgeCode + 3) / 4

		edgeStartValue := rectangles[group1[0]/4][vStart]

		for j := 1; j < group1Count; j++ {

			toCompare := rectangles[group1[0]/4][vStart]

			if toCompare < edgeStartValue {
				edgeStartValue = toCompare
				tmpIndex := group1[0]
				group1[0] = group1[j]
				group1[j] = tmpIndex
			}
		}

		edgeEndValue := rectangles[group1[0]/4][vEnd]

		for j := 1; j < group1Count; j++ {
			k := j
			for ; k < group1Count; k++ {
				toCompare := rectangles[group1[k]/4][vStart]

				if toCompare < edgeEndValue {
					return false
				}

				if toCompare == edgeEndValue {
					break
				}
			}

			if k == group1Count {
				return false
			}

			edgeEndValue = rectangles[group1[k]/4][vEnd]

			tmpIndex := group1[j]
			group1[j] = group1[k]
			group1[k] = tmpIndex
		}

		if group2Count == 0 {
			resultEdges[edgeCode] = true
			continue
		}

		edge2StartValue := rectangles[group2[0]/4][vStart]

		for j := 1; j < group2Count; j++ {

			toCompare := rectangles[group2[0]/4][vStart]

			if toCompare < edge2StartValue {
				edge2StartValue = toCompare
				tmpIndex := group2[0]
				group2[0] = group2[j]
				group2[j] = tmpIndex
			}
		}

		if edgeStartValue != edge2StartValue {
			return false
		}

		edge2EndValue := rectangles[group2[0]/4][vEnd]

		for j := 1; j < group2Count; j++ {
			k := j
			for ; k < group2Count; k++ {
				toCompare := rectangles[group2[k]/4][vStart]

				if toCompare < edge2EndValue {
					return false
				}

				if toCompare == edge2EndValue {
					break
				}
			}

			if k == group2Count {
				return false
			}

			edge2EndValue = rectangles[group2[k]/4][vEnd]

			tmpIndex := group2[j]
			group2[j] = group2[k]
			group2[k] = tmpIndex
		}

		if edgeEndValue != edge2EndValue {
			return false
		}
	}

	return true
}
