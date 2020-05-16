package leetcode

import (
	"testing"
)

func TestExample_1(t *testing.T) {
	nums := []int{-1, -1, 1}
	k := 0
	desire := 1

	output := subarraySum(nums, k)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_2(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	desire := 2

	output := subarraySum(nums, k)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}
