package leetcode

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	max := 1
	start := 0
	end := 1

	for ; end < len(s); end++ {
		for i := start; i < end; i++ {
			if s[i] == s[end] {
				if end-start > max {
					max = end - start
				}

				start = i + 1
			}
		}
	}

	if len(s)-start > max {
		return len(s) - start
	} else {
		return max
	}
}
