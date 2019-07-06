package sort

func Insertion(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			prev := j - 1
			if nums[j] < nums[prev] {
				nums[j], nums[prev] = nums[prev], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
