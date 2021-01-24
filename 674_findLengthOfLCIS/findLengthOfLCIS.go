package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currMaxLength := 0
	currLength := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currLength++
		} else {
			if currLength > currMaxLength {
				currMaxLength = currLength
			}

			currLength = 1
		}
	}

	if currLength > currMaxLength {
		return currLength
	} else {
		return currMaxLength
	}
}
