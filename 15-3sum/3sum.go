package _5_3sum

import "testing"

func Test3Sum(t *testing.T) {

}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result := []int{i, j, k}
					results = append(results, result)
				}
			}
		}
	}
	if len(results) != 0 {
		return results
	}
	return nil
}
