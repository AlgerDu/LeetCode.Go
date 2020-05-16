package leetcode

import "testing"

func TestExample(t *testing.T) {
	s := "abcabcbb"
	d := 3

	out := lengthOfLongestSubstring(s)

	if out != d {
		t.Errorf("out %d is not equal %d", out, d)
	}
}

func TestExample1(t *testing.T) {
	s := "bbbbb"
	d := 1

	out := lengthOfLongestSubstring(s)

	if out != d {
		t.Errorf("out %d is not equal %d", out, d)
	}
}

func TestExample2(t *testing.T) {
	s := "pwwkew"
	d := 3

	out := lengthOfLongestSubstring(s)

	if out != d {
		t.Errorf("out %d is not equal %d", out, d)
	}
}
