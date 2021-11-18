package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func Anaylse(arrayStr string) [][]int {

	arrayStr = strings.ReplaceAll(arrayStr, "[", "")
	arrayStr = strings.ReplaceAll(arrayStr, "]", "")

	numArray := strings.Split(arrayStr, ",")

	count := len(numArray)
	array := make([][]int, count/4)

	for i := 0; i < count; i++ {

		if i%4 == 0 {
			array[i/4] := make([]int, 4)
		}

		num, _ := strconv.Atoi(numArray[i])

		array[i/4][i%4] = num
	}

	return array
}

func TestExample_1(t *testing.T) {

	inputStr := "[[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]"

	desire := true

	result := isRectangleCover(Anaylse(inputStr))

	if desire != result {
		t.Errorf("result %t is not equal %t", result, desire)
	}
}
