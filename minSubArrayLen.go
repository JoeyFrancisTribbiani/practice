package main

func minSubArrayLen(target int, nums []int) int {
	var mix int = len(nums)
	len := len(nums)
	i, j := 0, 0
	sum := 0
	for i < len {
		for j < len {
			sum += nums[j]
			if sum < target {
				j++
			} else {
				break
			}
		}
		if i == 0 && j == len && sum < target {
			return 0
		}
		if sum < target {
			break
		}
		if mix > (j - i + 1) {
			mix = j - i + 1
		}
		for sum >= target {
			sum = sum - nums[i]
			i++
			if mix > (j-i+1) && sum >= target {
				mix = j - i + 1
			}
		}
		j++
	}
	return mix
}
