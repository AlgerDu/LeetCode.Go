package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func Anaylse(pointsStr string) [][]int {

	numArray := strings.Split(pointsStr, ",")

	point := make([][]int, len(numArray)/2)

	for i := 0; i < len(numArray); i++ {
		m := i / 2
		n := i % 2

		if point[m] == nil {
			point[m] = make([]int, 2)
		}

		num, _ := strconv.Atoi(numArray[i])

		point[m][n] = num
	}

	return point
}

func TestExample_01(t *testing.T) {
	pointsStr := "0,0,2,2,3,10,5,2,7,0"

	desire := 20

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}
