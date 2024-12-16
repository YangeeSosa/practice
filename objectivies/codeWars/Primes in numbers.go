package kata

import "strconv"

func primeFactors(n int) string {
	result := ""
	for i := 2; i*i <= n; i++ {
		count := 0
		for n%i == 0 {
			count++
			n /= i
		}
		if count > 0 {
			result += "(" + strconv.Itoa(i)
			if count > 1 {
				result += "**" + strconv.Itoa(count)
			}
			result += ")"
		}
	}
	if n > 1 {
		result += "(" + strconv.Itoa(n) + ")"
	}
	return result
}
