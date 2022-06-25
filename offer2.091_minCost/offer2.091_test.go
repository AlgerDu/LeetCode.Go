package leetcode

import (
	. "algerdu/leetcode/go/utils"
	"testing"
)

func TestExample_1(t *testing.T) {

	inputStr := "[[17,2,17],[16,16,5],[14,3,19]]"

	desire := 10

	result := minCost(CreateMatrix(inputStr, 3))

	if desire != result {
		t.Errorf("result %d is not equal %d", result, desire)
	}
}

func TestExample_2(t *testing.T) {

	inputStr := "[[7,6,2]]"

	desire := 2

	result := minCost(CreateMatrix(inputStr, 3))

	if desire != result {
		t.Errorf("result %d is not equal %d", result, desire)
	}
}
