package leetcode

import "testing"

func TestExample_1(t *testing.T) {

	desire := 2
	input := 4

	output := mySqrt(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_2(t *testing.T) {

	desire := 2
	input := 8

	output := mySqrt(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_18(t *testing.T) {

	desire := 2
	input := 4

	output := mySqrt(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}

func TestExample_610(t *testing.T) {

	desire := 2
	input := 5

	output := mySqrt(input)

	if output != desire {
		t.Errorf("output %d not equal %d", output, desire)
	}
}
