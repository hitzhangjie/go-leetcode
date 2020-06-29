package go_leetcode

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	locs := make(map[int][]int, len(nums))
	for idx, num := range nums {
		locs[num] = append(locs[num], idx)
	}
	fmt.Println("locs[n]posIdx:", locs)

	sort.Ints(nums)
	fmt.Println("sorted nums:", nums)

	stats := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			stats[0] = 0
			continue
		}
		if nums[i] == nums[i-1] { // ==
			stats[i] = stats[i-1]
		} else { // >
			stats[i] = i
		}
	}
	fmt.Println("stats to sorted nums:", stats)

	result := make([]int, len(nums))
	for posIdx, n := range stats {
		v := nums[posIdx]
		for _, loc := range locs[v] {
			result[loc] = n
		}
	}
	fmt.Println("result:", result)

	return result
}
