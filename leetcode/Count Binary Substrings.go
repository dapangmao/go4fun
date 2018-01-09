func countBinarySubstrings(s string) int {
    var dp = []int{1}
    var n = len(s)
    for i:=1; i<n; i++ {
        if s[i-1] != s[i] {
            dp = append(dp, 1)
        } else {
            dp[len(dp)-1]++
        }
    }
    var res = 0
    for i:=1; i<len(dp); i++ {
        if dp[i-1] < dp[i] {
            res += dp[i-1]
        } else {
            res += dp[i]
        }
    }
    return res
}
