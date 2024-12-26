package leetcode

func intToRoman(num int) string {
	result := ""
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, value := range values {
		for num >= value {
			num -= value
			result += symbols[i]
		}
	}
	return result
}
