package main

func removeDuplicatesK(nums []int, k int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	i, j, m := 0, 1, 1
	temp := nums[0]

	for j < len {
		m++
		if nums[j] == temp {
			if m <= k {
				i++
				nums[i] = nums[j]
			}
		} else {
			m = 1
			i++
			nums[i] = nums[j]
			temp = nums[i]
		}
		j++
	}

	return i + 1

}
