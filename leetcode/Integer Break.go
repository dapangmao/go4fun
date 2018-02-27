func integerBreak(n int) int {
    dp := make([]int, n+1)
    for i:=1; i<n+1; i++ {dp[i] = i}
    if n == 2 {return 1}
    if n == 3 {return 2}
    for i:=2; i<=n; i++ {
        for j:=0; j<i; j++ {
            var current = dp[j] * dp[i-j]
            if current > dp[i] {dp[i] = current} 
        }
    }
    return dp[n]
}
