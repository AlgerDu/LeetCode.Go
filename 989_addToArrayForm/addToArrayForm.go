package leetcode

func addToArrayForm(A []int, K int) []int {

	i := len(A) - 1

	for i >= 0 && K > 0 {
		K += A[i]

		A[i] = K % 10
		K = K / 10
		i--
	}

	if K == 0 {
		return A
	}

	j := 5
	tmp := make([]int, 6)

	for K > 0 {

		tmp[j] = K % 10

		K = K / 10
		j--
	}

	return append(tmp[j+1:6], A...)
}
