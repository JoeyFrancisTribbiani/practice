package stack

func dailyTemperatures(temperatures []int) []int {
	len := len(temperatures)
	res := make([]int, len)
	for i := len - 1; i < len; i-- {
		j := i - 1
		for j > 0 {
			if temperatures[i] > temperatures[j] {
				res[i-1] = i - j
				break
			}
			j--
		}
	}
	return res
}
