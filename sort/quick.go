package sort

func Quick(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}

func quick(nums []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := nums[start]
	left := start + 1
	right := end

	for ; left <= right; left++ {
		if pivot > nums[left] {
			continue
		}
		for ; left <= right; right-- {
			if pivot < nums[right] {
				continue
			}
			nums[left], nums[right] = nums[right], nums[left]
			break
		}
	}

	nums[right], nums[start] = pivot, nums[right]

	quick(nums, start, right-1)
	quick(nums, right+1, end)
}
