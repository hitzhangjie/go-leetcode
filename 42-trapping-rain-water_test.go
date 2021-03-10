package go_leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case-1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"case-2", args{[]int{4, 2, 0, 3, 2, 5}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack(t *testing.T) {
	stk := &stack{}

	stk.Push(1)
	assert.Equal(t, stk.Top(), 1)

	stk.Push(2)
	stk.Push(3)
	assert.Equal(t, stk.Empty(), false)

	stk.Pop()
	assert.Equal(t, stk.Top(), 2)

	stk.Pop()
	assert.Equal(t, stk.Top(), 1)

	stk.Pop()
	assert.Equal(t, stk.Empty(), true)

	assert.Equal(t, stk.Top(), int(math.MaxInt64))
}
