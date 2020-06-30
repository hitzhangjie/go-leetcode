package go_leetcode

import "testing"

func Test_findCols(t *testing.T) {
	type args struct {
		n    int
		rows int
	}
	tests := []struct {
		name     string
		args     args
		wantCols int
	}{
		{"case-1,1", args{1, 1}, 1},
		{"case-2,1", args{2, 1}, 2},
		{"case-3,2", args{3, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCols := findCols(tt.args.n, tt.args.rows); gotCols != tt.wantCols {
				t.Errorf("findCols() = %v, want %v", gotCols, tt.wantCols)
			}
		})
	}
}
