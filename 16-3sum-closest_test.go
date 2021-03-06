package go_leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case-1",
			args{
				[]int{-1, 2, 1, -4},
				1,
			},
			2,
		//}, {
		//	"case-2",
		//	args{
		//		[]int{0, 1, 2},
		//		3,
		//	},
		//	3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
