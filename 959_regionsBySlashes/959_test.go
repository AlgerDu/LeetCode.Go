package leetcode

import (
	"testing"
)

func TestExample_01(t *testing.T) {
	grid := []string{
		" /",
		"/ ",
	}

	desire := 2

	output := regionsBySlashes(grid)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_02(t *testing.T) {
	grid := []string{
		" /",
		"  ",
	}

	desire := 1

	output := regionsBySlashes(grid)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_03(t *testing.T) {
	grid := []string{
		"\\/",
		"/\\",
	}

	desire := 4

	output := regionsBySlashes(grid)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_04(t *testing.T) {
	grid := []string{
		"/\\",
		"\\/",
	}

	desire := 5

	output := regionsBySlashes(grid)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_05(t *testing.T) {
	grid := []string{
		"//",
		"/ ",
	}

	desire := 3

	output := regionsBySlashes(grid)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}

func TestExample_06(t *testing.T) {
	grid := []string{"/\\/\\", "\\///", " \\//", "////"}

	desire := 9

	output := regionsBySlashes(grid)

	if desire != output {
		t.Errorf("output %d is not equal %d", output, desire)
	}
}
