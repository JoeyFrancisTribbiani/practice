package set

import "sort"

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	hash := make(map[int]int)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			other := -nums[i]
			if v, ok := hash[other-nums[j]]; ok {
				if v > j {
					res = append(res, []int{nums[i], nums[j], other})
				}
			} else {
				break
			}
		}
	}
	return res
}
