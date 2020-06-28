package go_leetcode

import (
	"fmt"
	"math"
)

func checkParams(s int, nums []int) int {
	sum := 0
	for _, v := range nums {
		if v >= s {
			return 1
		}
		sum += v
	}
	if sum < s {
		return 0
	}
	return -1
}

func minSubArrayLen(s int, nums []int) int {

	// 先排除没有解的情况
	n := checkParams(s, nums)
	if n == 0 || n == 1 {
		return n
	}

	// 重点关注有解的情况
	minLength := math.MaxInt64
	minStat := stat{}
	stats := make([]stat, len(nums))

	for i := 0; i < len(nums); i++ {

		if i == 0 {
			stats[0].sum = nums[i]
			continue
		}

		st := stat{
			stats[i-1].startIdx,
			i,
			stats[i-1].sum + nums[i],
		}

		//step := 1

		if st.sum >= s {
			// 更新minStat
			if nLen := st.stopIdx - st.startIdx + 1; nLen <= minLength {
				minLength = nLen
				minStat = st
				// 要回溯一下看当前区间内去掉某些前缀，能否得到符合要求的更小长度的区间
				tmp := minStat.sum

				for j := minStat.startIdx; j < minStat.stopIdx; j++ {
					fmt.Printf("%d - %d = %d\n", tmp, nums[j], tmp-nums[j])
					tmp -= nums[j]
					if tmp >= s {
						minStat.startIdx++
						//step++
						minStat.sum = tmp
						minLength--
						//fmt.Println(minStat)
					} else {
						break
					}
				}
				st = minStat
			}

			for j := st.startIdx; j < st.startIdx+1; j++ {
				st.sum -= nums[j]
			}
			st.startIdx = st.startIdx + 1
			stats[i] = st

		} else {
			stats[i] = st
		}
	}

	return minLength
}

type stat struct {
	startIdx int
	stopIdx  int
	sum      int
}
