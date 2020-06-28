package go_leetcode

func xorOperation(n int, start int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = start + i*2
	}
	result := nums[0]
	for i := 0; i < n; i++ {
		result = nums[i]
	}
	return result
}
