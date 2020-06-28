package go_leetcode

import (
	"sort"
)

func threeSum(nums []int, target int) [][]int {

	results := [][]int{}

	// 首先要排下序，就变成有序数组r、l指针扫描问题了
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {

		// nums [i]> nums[i-1]过滤掉重复搜索的情况
		if i == 0 || nums[i] > nums[i-1] {

			for l, r := i+1, len(nums)-1; l < r; {

				if nums[i]+nums[l]+nums[r] == target {
					// 找到一个匹配，记录并调整搜索的左右边界
					results = append(results, []int{nums[i], nums[l], nums[r]})
					l++
					r--
					// 跳过相同的
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					// 跳过相同的
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else if nums[i]+nums[l]+nums[r] > target {
					// 调整右边界
					r--
				} else {
					// 调整左边界
					l++
				}
			}
		}
	}

	if len(results) != 0 {
		return results
	}
	return nil
}

// ThreeSumEqualTarget 从nums中找出三个数字相加之和等于target，排除重复的三元组
func ThreeSumEqualTarget(nums []int, target int) [][]int {
	return threeSum(nums, target)
}
