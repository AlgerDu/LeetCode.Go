package leetcode

import (
	"testing"
)

func TestExample_1(t *testing.T) {
	N := 10
	desire := 9

	output := monotoneIncreasingDigits(N)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_2(t *testing.T) {
	N := 1234
	desire := 1234

	output := monotoneIncreasingDigits(N)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_3(t *testing.T) {
	N := 332
	desire := 299

	output := monotoneIncreasingDigits(N)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}
