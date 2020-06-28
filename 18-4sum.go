package go_leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	results := [][]int{}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			t := target - nums[i]
			res := ThreeSumEqualTarget(nums[i+1:], t)
			if len(res) != 0 {
				for _, r := range res {
					v := []int{nums[i]}
					v = append(v, r...)
					results = append(results, v)
				}
			}
		}
	}

	if len(results) != 0 {
		return results
	}
	return nil
}

// FourSum 从nums中找出4个数字，它们的相加之和等于target，返回去重之后的四元组
func FourSum(nums []int, target int) [][]int {
	return fourSum(nums, target)
}
