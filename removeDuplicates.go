package main

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	i, j := 0, 1

	for j < len {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1

}
