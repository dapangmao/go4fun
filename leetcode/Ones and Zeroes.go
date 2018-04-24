func findMaxForm(strs []string, m int, n int) int {
    // 0-1 knapsack problem
    var dp = make([][]int, m+1)
    for i:=0; i<=m; i++ {dp[i] = make([]int, n+1)}
    for _, str := range strs {
        ones := strings.Count(str, "1") // count characters actually
        zeros := len(str) - ones
        for i:=m; i>=zeros; i-- {
            for j:=n; j>=ones; j-- {
                var current = dp[i-zeros][j-ones] + 1
                if current > dp[i][j] {
                    dp[i][j] = current
                }
            }
        }
    }
    return dp[m][n]
}
