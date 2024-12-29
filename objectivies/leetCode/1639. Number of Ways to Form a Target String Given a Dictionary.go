package leetcode

const MOD = 1e9 + 7

func numWays(words []string, target string) int {
	m := len(words[0])
	n := len(target)

	freq := make([][]int, m)
	for i := range freq {
		freq[i] = make([]int, 26)
	}

	for _, word := range words {
		for j := 0; j < m; j++ {
			freq[j][word[j]-'a']++
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1

	for i := 0; i <= n; i++ {
		for j := 0; j < m; j++ {
			if i < n {
				dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]*freq[j][target[i]-'a']) % MOD
			}
			dp[i][j+1] = (dp[i][j+1] + dp[i][j]) % MOD
		}
	}

	return dp[n][m]
}
