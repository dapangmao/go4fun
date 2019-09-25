func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    var dp = make([]int, n)
    copy(dp, cost)
    for i:=2; i<n; i++ {
        if dp[i-2] < dp[i-1] {
            dp[i] += dp[i-2]
        } else {
            dp[i] += dp[i-1]
        }
    }
    if dp[n-1] < dp[n-2] {return dp[n-1]}
    return dp[n-2]
}
