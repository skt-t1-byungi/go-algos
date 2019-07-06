package sort

func Selection(list []int) []int {
	for n := range list {
		i := -1
		for j := n; j < len(list); j++ {
			if v := list[j]; i == -1 || v < list[i] {
				i = j
			}
		}
		list[n], list[i] = list[i], list[n]
	}
	return list
}
