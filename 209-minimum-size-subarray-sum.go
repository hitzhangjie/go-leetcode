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

	// 先排除特殊解的情况
	n := checkParams(s, nums)
	if n == 0 || n == 1 {
		return n
	}
	// 重点关注有解的情况
	minlen := math.MaxInt64
	minstat := stat{}
	preStat := stat{}
	curStat := stat{}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			curStat.sum = nums[i]
			continue
		}
		preStat = curStat

		st := stat{
			preStat.startIdx,
			i,
			preStat.sum + nums[i],
		}
		if st.sum < s {
			curStat = st
			continue
		}

		step := 1
		if nlen := st.stopIdx - st.startIdx + 1; nlen > minlen {
			// todo 优化查询效率的问题，不应该影响正确性
			//step += nlen - minlen
			step = 1
			goto HERE

		} else {
			minlen = nlen
			minstat = st
			// 当前区间去掉某些前缀，能否得到符合要求的更小长度的区间
			tmp := minstat.sum
			for j := minstat.startIdx; j < minstat.stopIdx; j++ {
				//fmt.Printf("%d - %d = %d\n", tmp, nums[j], tmp-nums[j])
				tmp -= nums[j]
				if tmp >= s {
					minstat.startIdx++
					minstat.sum = tmp
					minlen--
					//fmt.Println(minstat)
				} else {
					break
				}
			}
			st = minstat
		}
	HERE:
		for j, jm := st.startIdx, st.startIdx+step; j < jm && j < st.stopIdx; j++ {
			st.startIdx++
			st.sum -= nums[j]
		}
		curStat = st
	}
	fmt.Println(minstat)

	return minlen
}

type stat struct {
	startIdx int
	stopIdx  int
	sum      int
}
