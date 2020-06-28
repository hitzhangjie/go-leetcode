package go_leetcode

import (
	"bytes"
)

var matchesParenthese = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func generateParenthesis(n int) []string {
	result := []string{}

	stack := NewStack()

	push1 := func() {
		stack.Push('(')
	}
	push2 := func() {
		stack.Push(')')
	}

	// 6!=720种排列组合，这个测试是ok的
	fns := []func(){}
	for i := 0; i < n; i++ {
		fns = append(fns, push1)
		//fns = append(fns, push2)
	}
	for i := 0; i < n; i++ {
		//fns = append(fns, push1)
		fns = append(fns, push2)
	}
	fnsAll := permute(fns)

	existed := map[string]struct{}{}
	for _, fns := range fnsAll {
		for _, fn := range fns {
			fn()
		}
		//fmt.Printf("%d stack:%s\n", idx+1, stack.Inspect())
		if desc, ok := isValidParenthese(stack); ok && len(desc) != 0 {
			if _, ok := existed[desc]; ok {
				continue
			}
			result = append(result, desc)
			existed[desc] = struct{}{}
		}
		stack.Clear()
	}
	return result
}

func isValidParenthese(stack *Stack) (desc string, valid bool) {

	if stack.IsEmpty() {
		return "", true
	}

	stk := NewStack()
	for stack.Len() > 0 {
		v := stack.Pop()
		stk.Push(v)
	}

	buf := bytes.Buffer{}
	for stk.Len() > 0 {
		v := stk.Pop()
		buf.WriteString(string([]rune{v.(rune)}))
		if stack.IsEmpty() {
			stack.Push(v)
			continue
		}
		el := stack.Peak()
		if matchesParenthese[el.(rune)] == v.(rune) {
			stack.Pop()
		} else {
			stack.Push(v)
		}
	}
	if stack.IsEmpty() {
		return buf.String(), true
	}
	return "", false
}

// permute 排列组合
func permute(fns []func()) [][]func() {
	if len(fns) == 1 {
		return [][]func(){fns}
	}
	if len(fns) == 2 {
		return [][]func(){
			{fns[0], fns[1]},
			{fns[1], fns[0]},
		}
	}

	var fnsAll = [][]func(){}
	pms := permute(fns[1:])
	for _, pm := range pms {

		for idx, _ := range pm {
			v := []func(){}
			v = append(v, pm[0:idx]...)
			v = append(v, fns[0])
			v = append(v, pm[idx:]...)
			fnsAll = append(fnsAll, v)
		}
		v := []func(){}
		v = append(v, pm...)
		v = append(v, fns[0])
		fnsAll = append(fnsAll, v)
	}
	return fnsAll
}
