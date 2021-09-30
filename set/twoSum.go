package set

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if v, ok := hash[other]; ok {
			return []int{v, i}
		}
		hash[nums[i]] = i
	}
	return []int{0, 1}
}
