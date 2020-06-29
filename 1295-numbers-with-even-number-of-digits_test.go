package go_leetcode

import "testing"

func Test_hasEvenNumberOfDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case-1", args{1}, false},
		{"case-12", args{12}, true},
		{"case-345", args{345}, false},
		{"case-2", args{2}, false},
		{"case-6", args{6}, false},
		{"case-7896", args{7896}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasEvenNumberOfDigits(tt.args.num); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case-[12,345,2,6,7896]", args{[]int{12, 345, 2, 6, 7896}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
