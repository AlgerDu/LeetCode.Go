package leetcode

import (
	"testing"
)

func TestExample_1(t *testing.T) {

	pattern := "abba"
	str := "dog cat cat dog"
	desire := true

	output := wordPattern(pattern, str)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_2(t *testing.T) {

	pattern := "abba"
	str := "dog cat cat fish"
	desire := false

	output := wordPattern(pattern, str)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_3(t *testing.T) {

	pattern := "aaaa"
	str := "dog cat cat dog"
	desire := false

	output := wordPattern(pattern, str)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_5(t *testing.T) {

	pattern := "abba"
	str := "dog dog dog dog"
	desire := false

	output := wordPattern(pattern, str)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_6(t *testing.T) {

	pattern := "aaa"
	str := "aa aa aa aa"
	desire := false

	output := wordPattern(pattern, str)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}
