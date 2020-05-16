package leetcode

func subarraySum(nums []int, k int) int {
	count := 0
	i := -1
	j := 0
	total := 0

	for ; j < len(nums); j++ {

		if nums[j] == k {
			count++

			i = j
			total = 0

		} else {
			total += nums[j] + 1000

			if total == (k + 1000*(j-i)) {
				count++
			}

			for total >= (k+1000*(j-i)) && i < j {
				total -= nums[i+1] + 1000
				i++
			}
		}
	}

	return count
}
