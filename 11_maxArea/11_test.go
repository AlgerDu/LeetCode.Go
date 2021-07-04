package leetcode

import "testing"

func TestExample_1(t *testing.T) {

	desire := 49
	input := []int{1,8,6,2,5,4,8,3,7}

	output := maxArea(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_2(t *testing.T) {

	desire := 1
	input := []int{1,1}

	output := maxArea(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_3(t *testing.T) {

	desire := 16
	input := []int{4,3,2,1,4}

	output := maxArea(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_4(t *testing.T) {

	desire := 2
	input := []int{1,2,1}

	output := maxArea(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_5(t *testing.T) {

	desire := 8
	input := []int{1,0,0,0,0,0,0,2,2}

	output := maxArea(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}