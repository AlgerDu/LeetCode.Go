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

func AnaylseToString(array []int) string {

	str := "[" + strconv.Itoa(array[0])

	for i := 1; i < len(array); i++ {
		str += "," + strconv.Itoa(array[i])
	}

	return str + "]"
}

func TestExample_01(t *testing.T) {
	aStr := "[1,2,0,0]"
	k := 34

	desire := "[1,2,3,4]"

	output := AnaylseToString(addToArrayForm(Anaylse(aStr), k))

	if desire != output {
		t.Errorf("output %s is not equal %s", output, desire)
	}
}

func TestExample_02(t *testing.T) {
	aStr := "[2,7,4]"
	k := 181

	desire := "[4,5,5]"

	output := AnaylseToString(addToArrayForm(Anaylse(aStr), k))

	if desire != output {
		t.Errorf("output %s is not equal %s", output, desire)
	}
}

func TestExample_03(t *testing.T) {
	aStr := "[9,9,9,9,9,9,9,9,9,9]"
	k := 1

	desire := "[1,0,0,0,0,0,0,0,0,0,0]"

	output := AnaylseToString(addToArrayForm(Anaylse(aStr), k))

	if desire != output {
		t.Errorf("output %s is not equal %s", output, desire)
	}
}
