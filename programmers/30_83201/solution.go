package programmers_30_83201

func solution(scores [][]int) string {
	ret := ""
	for i := 0; i < len(scores); i++ {
		max, min, sum := 0, 100, 0
		for j := 0; j < len(scores); j++ {
			if i == j {
				continue
			}
			val := scores[j][i]
			sum += val
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
		self := scores[i][i]
		avg := 0
		if self > max || self < min {
			avg = sum / (len(scores) - 1)
		} else {
			avg = (sum + self) / len(scores)
		}
		switch {
		case avg >= 90:
			ret += "A"
		case avg >= 80:
			ret += "B"
		case avg >= 70:
			ret += "C"
		case avg >= 50:
			ret += "D"
		default:
			ret += "F"
		}
	}
	return ret
}