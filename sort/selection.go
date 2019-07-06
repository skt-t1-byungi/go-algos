package sort

func Selection(nums []int) []int {
	for n := range nums {
		i := n
		for j := n + 1; j < len(nums); j++ {
			if v := nums[j]; v < nums[i] {
				i = j
			}
		}
		nums[n], nums[i] = nums[i], nums[n]
	}
	return nums
}
