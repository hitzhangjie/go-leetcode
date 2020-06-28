package go_leetcode

import (
	"bytes"
	"fmt"
)

var matches = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	stack := NewStack()
	for _, c := range s {
		if stack.IsEmpty() {
			stack.Push(c)
			continue
		}
		el := stack.Peak()
		if matches[el.(rune)] == c {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

type Node struct {
	val interface{}
}

type Stack struct {
	els []Node
}

func NewStack() *Stack {
	return &Stack{
		els: make([]Node, 0, 16),
	}
}

func (s *Stack) Push(val interface{}) {
	s.els = append(s.els, Node{val})
}

func (s *Stack) Pop() interface{} {
	n := len(s.els)
	if n == 0 {
		panic("empty stack")
	}
	val := s.els[n-1]
	s.els = s.els[:n-1]
	return val.val
}

func (s *Stack) Peak() interface{} {
	if n := len(s.els); n != 0 {
		return s.els[n-1].val
	} else {
		panic("empty stack")
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.els) == 0
}

func (s *Stack) Len() int {
	return len(s.els)
}

func (s *Stack) Inspect() string {
	buf := bytes.Buffer{}
	for _, el := range s.els {
		buf.WriteString(fmt.Sprintf("%c", el.val))
	}
	return buf.String()
}
func (s *Stack) Clear() {
	s.els = s.els[0:0]
}
