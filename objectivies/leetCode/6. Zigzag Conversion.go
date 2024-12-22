package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	for _, v := range s {
		rows[curRow] += string(v)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}
	return strings.Join(rows, "")
}
