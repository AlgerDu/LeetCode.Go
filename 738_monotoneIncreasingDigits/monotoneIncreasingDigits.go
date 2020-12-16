package leetcode

func monotoneIncreasingDigits(N int) int {

	i := 10
	offset := N % 10
	N = N / 10
	max := offset

	for N > 0 {
		tmpOffset := N % 10
		N = N / 10

		if tmpOffset <= offset {
			max = max + tmpOffset*i
			offset = tmpOffset
		} else {
			offset = tmpOffset - 1
			max = offset*i + i - 1
		}

		i = i * 10
	}

	return max
}
