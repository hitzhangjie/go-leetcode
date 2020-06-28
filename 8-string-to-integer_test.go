package go_leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case-1", args{"42"}, 42},
		{"case-2", args{"   -42"}, -42},
		{"case-3", args{"4193 with words"}, 4193},
		{"case-4", args{"9223372036854775808"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
