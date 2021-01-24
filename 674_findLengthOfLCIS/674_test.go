package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func Anaylse(arrayStr string) []int {

	arrayStr = strings.ReplaceAll(arrayStr, "[", "")
	arrayStr = strings.ReplaceAll(arrayStr, "]", "")

	numArray := strings.Split(arrayStr, ",")

	array := make([]int, len(numArray))

	for i := 0; i < len(numArray); i++ {

		num, _ := strconv.Atoi(numArray[i])

		array[i] = num
	}

	return array
}

func TestExample_01(t *testing.T) {
	aStr := "[1,3,5,4,7]"

	desire := 3

	output := findLengthOfLCIS(Anaylse(aStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_02(t *testing.T) {
	aStr := "[2,2,2,2,2]"

	desire := 1

	output := findLengthOfLCIS(Anaylse(aStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_03(t *testing.T) {
	aStr := "[2]"

	desire := 1

	output := findLengthOfLCIS(Anaylse(aStr))

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}
