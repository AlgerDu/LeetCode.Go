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
			array[i/4] = make([]int, 4)
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

func TestExample_2(t *testing.T) {

	inputStr := "[[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]"

	desire := false

	result := isRectangleCover(Anaylse(inputStr))

	if desire != result {
		t.Errorf("result %t is not equal %t", result, desire)
	}
}

func TestExample_3(t *testing.T) {

	inputStr := "[[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]"

	desire := false

	result := isRectangleCover(Anaylse(inputStr))

	if desire != result {
		t.Errorf("result %t is not equal %t", result, desire)
	}
}

func TestExample_4(t *testing.T) {

	inputStr := "[[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]"

	desire := false

	result := isRectangleCover(Anaylse(inputStr))

	if desire != result {
		t.Errorf("result %t is not equal %t", result, desire)
	}
}

func TestExample_41(t *testing.T) {

	inputStr := "[[0,0,4,1],[7,0,8,2],[6,2,8,3],[5,1,6,3],[4,0,5,1],[6,0,7,2],[4,2,5,3],[2,1,4,3],[0,1,2,2],[0,2,2,3],[4,1,5,2],[5,0,6,1]]"

	desire := true

	result := isRectangleCover(Anaylse(inputStr))

	if desire != result {
		t.Errorf("result %t is not equal %t", result, desire)
	}
}
