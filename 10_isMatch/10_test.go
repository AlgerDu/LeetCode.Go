package leetcode

import (
	"testing"
)

// 输入：s = "aa" p = "a"
// 输出：false
// 解释："a" 无法匹配 "aa" 整个字符串。

func TestExample_01(t *testing.T) {
	s := "aa"
	p := "a"

	desire := false

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

// 输入：s = "aa" p = "a*"
// 输出：true
// 解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

func TestExample_02(t *testing.T) {
	s := "aa"
	p := "a*"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

// 输入：s = "ab" p = ".*"
// 输出：true
// 解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

func TestExample_03(t *testing.T) {
	s := "ab"
	p := ".*"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

// 输入：s = "aab" p = "c*a*b"
// 输出：true
// 解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

func TestExample_04(t *testing.T) {
	s := "aab"
	p := "c*a*b"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

// 输入：s = "mississippi" p = "mis*is*p*."
// 输出：false

func TestExample_05(t *testing.T) {
	s := "mississippi"
	p := "mis*is*p*."

	desire := false

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

// "mississippi"
// "mis*is*ip*."
// true

func TestExample_06(t *testing.T) {
	s := "mississippi"
	p := "mis*is*ip*."

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_07(t *testing.T) {
	s := "ab"
	p := ".*c"

	desire := false

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_08(t *testing.T) {
	s := "aaa"
	p := "aaaa"

	desire := false

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_09(t *testing.T) {
	s := "aaa"
	p := "a*a"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_10(t *testing.T) {
	s := "aaab"
	p := "a*ab"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_11(t *testing.T) {
	s := "aaa"
	p := "ab*a*c*a"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_12(t *testing.T) {
	s := "aaaaac"
	p := "a*.c"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_13(t *testing.T) {
	s := "aaaabc"
	p := "a*.c"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_14(t *testing.T) {
	s := "ac"
	p := "a*.c"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_15(t *testing.T) {
	s := "c"
	p := "a*.c"

	desire := false

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_16(t *testing.T) {
	s := "aaaaaac"
	p := "a*.ac"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}

func TestExample_17(t *testing.T) {
	s := "aaaaaac"
	p := "a*.aac"

	desire := true

	output := isMatch(s, p)

	if desire != output {
		t.Errorf("output %t is not equal %t", output, desire)
	}
}
