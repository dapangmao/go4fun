func minDistance(word1 string, word2 string) int {
    n, m := len(word1), len(word2)
    var dp = make([][]int, n+1)
    for i:=0; i<=n; i++ {dp[i] = make([]int, m+1)}
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            current := dp[i][j]
            if word1[i] == word2[j] {current++}
            dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1], current)
        }
    }
    return n + m - 2*dp[n][m]
}


func max(a, b, c int) int {
    res := a
    if b > res {res = b}
    if c > res {res = c}
    return c
}
