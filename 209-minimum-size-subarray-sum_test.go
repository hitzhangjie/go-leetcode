package go_leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case-1",
			args{
				7,
				[]int{2, 3, 1, 2, 4, 3},
			},
			2,
		}, {
			"case-2",
			args{
				11,
				[]int{1, 2, 3, 4, 5},
			},
			3,
		}, {
			"case-3",
			args{
				15,
				[]int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			},
			2,
		}, {
			"case-4",
			args{
				80,
				[]int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
