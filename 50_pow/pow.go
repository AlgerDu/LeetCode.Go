package leetcode

func myPow(x float64, n int) float64 {

	if n < 0 {
		x = 1 / x
		n = n * -1
	}

	curr := 1
	total := x

	for curr*2 <= n {
		total *= total
		curr *= 2
	}

	if curr == n {
		return total
	}

	if n-curr < curr*2-n {
		return total * myPow(x, n-curr)
	} else {
		return total * total / myPow(x, curr*2-n)
	}
}
