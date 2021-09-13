package programmers_30_86048

func solution(enter []int, leave []int) []int {
	i := 0
	room := make(map[int]bool)
	ret := make([]int, len(enter))
	for _, n := range leave {
		for ; !room[n]; i++ {
			m := enter[i]
			for k := range room {
				ret[k-1]++
			}
			ret[m-1] = len(room)
			room[m] = true
		}
		delete(room, n)
	}
	return ret
}
