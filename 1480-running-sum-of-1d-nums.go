package go_leetcode

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	sums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sums[i] = nums[i]
		} else {
			sums[i] = nums[i] + sums[i-1]
		}
	}
	return sums
}
