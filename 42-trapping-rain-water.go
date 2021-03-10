package go_leetcode

import "math"

// 有点类似递减栈，只是修改成：
// 指定数组T[]，对T[]中从第一个不为0的元素开始，寻找不小于T[i]的第一个T[j]，攒的水就是j-i-1
func trap(height []int) int {
	stk := &stack{}
	water := 0

	skip := true

	for i, h := range height {
		if skip && h == 0 {
			skip = false
			continue
		}

		for !stk.Empty() && h >= height[stk.Top()] {
			top := stk.Top()
			water += i - top - 1
			stk.Pop()
		}
		stk.Push(i)
	}
	return water
}

type stack struct {
	elements []int
}

func (s *stack) Push(x int) {
	s.elements = append(s.elements, x)
}

func (s *stack) Pop() {
	if s.elements != nil {
		n := len(s.elements)
		s.elements = s.elements[:n-1]
	}
}

func (s *stack) Top() int {
	n := len(s.elements)
	if n != 0 {
		return s.elements[n-1]
	}
	return int(math.MaxInt64)
}

func (s *stack) Empty() bool {
	return len(s.elements) == 0
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
