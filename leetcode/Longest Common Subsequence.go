func longestCommonSubsequence(text1 string, text2 string) int {
    n, m := len(text1), len(text2)
    if n == 0 || m == 0 {return 0}
    dp := make([][]int, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = make([]int, m+1)
    }
    for i:=1; i<=n; i++ {
        for j:=1; j<=m; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = dp[i-1][j] 
                if dp[i][j-1] > dp[i][j] {dp[i][j] = dp[i][j-1]}
            }
        }
    }
    return dp[n][m]
}
