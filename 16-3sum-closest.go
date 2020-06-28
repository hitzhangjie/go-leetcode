package go_leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	var (
		val     int = math.MaxInt64
		minDiff int = math.MaxInt64
	)

	for i := 0; i < len(nums); i++ {

		for l, r := i+1, len(nums)-1; l < r; {

			sum := nums[i] + nums[l] + nums[r]
			diff := sum - target
			if diff < 0 {
				diff = -1 * diff
			}

			if diff < minDiff {
				minDiff = diff
				val = sum
			}

			if sum == target {
				return val
			} else if sum > target {
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}

		}
	}
	return val
}
