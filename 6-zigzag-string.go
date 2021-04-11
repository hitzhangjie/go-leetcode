package go_leetcode

func convert(s string, numRows int) string {
	numCols := findCols(len(s), numRows)

	// 矩阵大小 numRows x numCols，列式存储就可以了
	matrix := newMatrx(numRows, numCols)
	_ = matrix

	it := 0
	_ = it
	for cIdx := 0; cIdx < numRows; cIdx++ {
		topToBottom := true
		_ = topToBottom
		if cIdx == 0 {

		} else {
		}

		for rIdx := 0; rIdx < numCols; rIdx++ {
		}
	}

	return ""
}

// 创建一个矩阵
func newMatrx(numRows, numCols int) [][]rune {

	matrix := make([][]rune, numRows, numRows)
	for i := 0; i < numRows; i++ {
		matrix[0] = []rune{}
		for j := 0; j < numCols; j++ {
			matrix[0][j] = 0
		}
	}
	return matrix
}

// 存储n个元素，指定rows，找到最合适的cols，构成一个矩阵rowsxcols
func findCols(n int, rows int) (cols int) {
	d := n / rows
	e := n % rows
	if e == 0 {
		return d
	}
	return d + 1
}
