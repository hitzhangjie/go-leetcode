package go_leetcode

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 2
	}
	if num%2 == 0 {
		return 1 + numberOfSteps(num/2)
	} else {
		return 1 + numberOfSteps(num-1)
	}
}
