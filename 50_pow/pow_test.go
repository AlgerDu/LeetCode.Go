package leetcode

import "testing"

func TestExample(t *testing.T) {
	x := 2.00000
	n := 10

	desire := 1024.0000000

	output := myPow(x, n)

	if output != desire {
		t.Errorf("output %f is not equal %f", output, desire)
	}
}

func TestExample2(t *testing.T) {
	x := 2.10000
	n := 3

	desire := 9.26100

	output := myPow(x, n)

	if output != desire {
		t.Errorf("output %f is not equal %f", output, desire)
	}
}

func TestExample3(t *testing.T) {
	x := 2.00000
	n := -2

	desire := 0.25000

	output := myPow(x, n)

	if output != desire {
		t.Errorf("output %f is not equal %f", output, desire)
	}
}
