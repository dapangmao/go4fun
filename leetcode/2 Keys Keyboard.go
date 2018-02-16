func minSteps(n int) int {
    dp := make([]int, n+1)
    for i:=2; i<n+1; i++ {
        dp[i] = i
    }
    for i:=4; i<n+1; i++ {
        for j:=2; j<i/2+1; j++ {
            if i % j != 0 {continue}
            var current = dp[j] + i / j
            if current < dp[i] {dp[i] = current}
        }
    }
    return dp[n]
}
