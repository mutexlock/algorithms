package backtrace

import (
	"bytes"
)

func solveNQueens(n int) [][]string {
	var res [][]string
	if n == 0 {
		return res
	}
	dfs(&res, []int{}, n)
	return res
}

func dfs(res *[][]string, cols []int, n int) {
	if n == len(cols) {
		*res = append(*res, draw(cols, n))
		return
	}
	for i := 0; i < n; i++ {
		if !isOK(cols, i) {
			continue
		}
		cols = append(cols, i)
		dfs(res, cols, n)
		cols = cols[:len(cols)-1]
	}
	return
}

func isOK(cols []int, coln int) bool {
	row := len(cols)
	for i := 0; i < row; i++ {
		if cols[i] == coln {
			return false
		}
		if cols[i] == coln+row-i {
			return false
		}
		if cols[i] == coln-row+i {
			return false
		}
	}
	return true
}

func draw(cols []int, n int) []string {
	var cb []string
	for i := 0; i < n; i++ {
		var buffer bytes.Buffer
		for j := 0; j < n; j++ {
			if cols[i] == j {
				buffer.WriteString("Q")
			} else {
				buffer.WriteString(".")
			}
		}
		cb = append(cb, buffer.String())
	}
	return cb
}
