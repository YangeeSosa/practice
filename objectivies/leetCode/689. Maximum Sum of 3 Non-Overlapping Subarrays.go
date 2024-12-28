package leetcode

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	sum_k := make([]int, n-k+1)
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		if i >= k {
			sum -= nums[i-k]
		}
		if i >= k-1 {
			sum_k[i-k+1] = sum
		}
	}
	left := make([]int, n-k+1)
	right := make([]int, n-k+1)

	best := 0
	for i := 0; i < len(sum_k); i++ {
		if sum_k[i] > sum_k[best] {
			best = i
		}
		left[i] = best
	}

	best = len(sum_k) - 1
	for i := len(sum_k) - 1; i >= 0; i-- {
		if sum_k[i] >= sum_k[best] {
			best = i
		}
		right[i] = best
	}

	result := []int{-1, -1, -1}
	max_sum := 0
	for j := k; j < len(sum_k)-k; j++ {
		i, l := left[j-k], right[j+k]
		total := sum_k[i] + sum_k[j] + sum_k[l]
		if total > max_sum || (total == max_sum && (result[0] == -1 || i < result[0] || (i == result[0] && j < result[1]))) {
			max_sum = total
			result = []int{i, j, l}
		}
	}
	return result
}
