package stack

func FindTargetSumWays(nums []int, target int) int {
	count := 0
	count = dfs(nums, 0, 0, target, count)
	return count
}
func dfs(nums []int, i int, sum int, target int, count int) int {
	len := len(nums)
	if i > len-1 {
		if sum == target {
			return count + 1
		}
		return count
	}
	count = dfs(nums, i+1, sum+nums[i], target, count)
	count = dfs(nums, i+1, sum-nums[i], target, count)
	return count
}
