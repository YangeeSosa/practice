package kata

import "sort"

func ExpressionMatter(a int, b int, c int) int {
	arr := []int{a * (b + c), a * b * c, a + b + c, (a + b) * c}
	sort.Ints(arr)
	return arr[len(arr)-1]
}
