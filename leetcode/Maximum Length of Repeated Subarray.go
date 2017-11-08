func findLength(A []int, B []int) int {
    n, m := len(A), len(B)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    } 
    for i:=n-1; i>=0; i-- {
        for j:=m-1; j>=0; j-- {
            if A[i] == B[j] {
                dp[i][j] = dp[i+1][j+1] + 1
            }
        }
    }
    res := -1
    for _, x := range dp {
        for _, y := range x {
            if y > res {
                res = y
            }
        }
    }
    return res
}
