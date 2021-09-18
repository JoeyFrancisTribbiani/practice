package main

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
		}
		if numbers[i]+numbers[j] < target {
			i++
		}
		if numbers[i]+numbers[j] == target {
			break
		}
	}
	return []int{i + 1, j + 1}
}
