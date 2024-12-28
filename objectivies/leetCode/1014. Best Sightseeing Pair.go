package leetcode

func maxScoreSightseeingPair(values []int) int {
	result := 0
	maxI := values[0]
	for i := 1; i < len(values); i++ {
		result = max(result, maxI+values[i]-i)
		maxI = max(maxI, values[i]+i)
	}
	return result
}
