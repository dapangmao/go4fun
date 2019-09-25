func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][m] = dp[i+1][m] + int(s1[i])
	}
	for i := m - 1; i >= 0; i-- {
		dp[n][i] = dp[n][i+1] + int(s2[i])
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				opt1, opt2 := dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j])
				if opt1 >= opt2 {
					dp[i][j] = opt2
				} else {
					dp[i][j] = opt1
				}
			}
		}
	}
	return dp[0][0]
}
