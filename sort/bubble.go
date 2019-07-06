package sort

func Bubble(nums []int) []int {
	for n := range nums {
		for i := 0; i < len(nums)-(n+1); i++ {
			j := i + 1
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
