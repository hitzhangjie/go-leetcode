package go_leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case-1",
			args{[]int{1, 0, -1, 0, -2, 2}, 0},
			[][]int{
				{-1, 0, 0, 1},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, 0)

			m1 := map[string]struct{}{}
			for _, v := range got {
				m1[fmt.Sprintf("%d%d%d%d", v[0], v[1], v[2], v[3])] = struct{}{}
			}

			m2 := map[string]struct{}{}
			for _, v := range tt.want {
				m2[fmt.Sprintf("%d%d%d%d", v[0], v[1], v[2], v[3])] = struct{}{}
			}

			if !reflect.DeepEqual(m1, m2) {
				t.Errorf("fourSum() = %v, want %v", m1, m2)
			}
		})
	}
}
