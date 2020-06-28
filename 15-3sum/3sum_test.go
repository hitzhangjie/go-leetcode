package _5_3sum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case-1",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)

			m1 := map[string]struct{}{}
			for _, v := range got {
				m1[fmt.Sprintf("%d%d%d", v[0], v[1], v[2])] = struct{}{}
			}

			m2 := map[string]struct{}{}
			for _, v := range tt.want {
				m2[fmt.Sprintf("%d%d%d", v[0], v[1], v[2])] = struct{}{}
			}

			if !reflect.DeepEqual(m1, m2) {
				t.Errorf("threeSum() = %v, want %v", m1, m2)
			}
		})
	}
}
