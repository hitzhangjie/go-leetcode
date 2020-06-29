package go_leetcode

func findNumbers(nums []int) int {
	cnt := 0
	for _, num := range nums {
		if hasEvenNumberOfDigits(num) {
			cnt++
		}
	}
	return cnt
}

func hasEvenNumberOfDigits(num int) bool {
	cnt := 0
	for {
		n := num / 10
		cnt++

		if n == 0 {
			break
		}
		num = n
	}

	if cnt%2 == 0 {
		return true
	}
	return false
}
