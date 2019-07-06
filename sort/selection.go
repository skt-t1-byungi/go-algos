package sort

func Selection(list []int) []int {
	for n := range list {
		i := n
		for j := n + 1; j < len(list); j++ {
			if v := list[j]; v < list[i] {
				i = j
			}
		}
		list[n], list[i] = list[i], list[n]
	}
	return list
}
