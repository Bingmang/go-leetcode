package _494_target_sum

func findTargetSumWays(nums []int, S int) int {
	return findTargetSumWaysCore(nums, 0, S)
}

func findTargetSumWaysCore(nums []int, start int, S int) int {
	if start == len(nums) {
		if S == 0 {
			return 1
		} else {
			return 0
		}
	}
	return findTargetSumWaysCore(nums, start + 1, S + nums[start]) + findTargetSumWaysCore(nums, start + 1, S - nums[start])
}
