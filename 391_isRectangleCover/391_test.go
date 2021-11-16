package leetcode

import (
	"testing"
)

func TestExample_1(t *testing.T) {

	desire := true

	result := isRectangleCover(make([][]int, 1))

	if desire != result {
		t.Errorf("result %t is not equal %t", result, desire)
	}
}
