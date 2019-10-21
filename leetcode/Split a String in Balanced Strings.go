func balancedStringSplit(s string) int {
    n := len(s)
    dp := make([]int, n+1)
    for i:=2; i<=n; i++ {
        for j:= 0; j<i-1; j++ {
            if (i-j) % 2 == 0 && strings.Count(s[j:i], "R") == (i-j) / 2 {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    return dp[n]
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
