package leetcode

func mySqrt(x int) int {

	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	if x == 2 {
		return 1
	}

	s := 1
	e := x

	for e-s > 1 {
		m := (s + e) / 2

		if m*m == x {
			return m
		}

		if m*m < x {
			s = m
		} else {
			e = m
		}
	}

	return s
}
