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

func TestExample_02(t *testing.T) {
	pointsStr := "3,12,-2,5,-4,1"

	desire := 18

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_03(t *testing.T) {
	pointsStr := "0,0,1,1,1,0,-1,1"

	desire := 4

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_04(t *testing.T) {
	pointsStr := "-1000000,-1000000,1000000,1000000"

	desire := 4000000

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_05(t *testing.T) {
	pointsStr := "0,0"

	desire := 0

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_06(t *testing.T) {
	pointsStr := "0,0,2,2,3,10,15,2,17,0"

	desire := 30

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_07(t *testing.T) {
	pointsStr := "-14,-14,-18,5,18,-10,18,18,10,-2"

	desire := 102

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_08(t *testing.T) {
	pointsStr := "5,8,18,-6,-18,13,-8,-13,-13,3,-15,2,-12,17,14,16,-4,3,-17,-7,8,9,17,14,-13,2,-3,-1,4,-20"

	desire := 143

	output := minCostConnectPoints(Anaylse(pointsStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}
