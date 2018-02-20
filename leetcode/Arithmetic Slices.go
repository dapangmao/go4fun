func numberOfArithmeticSlices(A []int) int {
    sum := 0
    dp := make([]int, len(A))
    for i:=2; i<len(A); i++ {
        if A[i] - A[i-1] == A[i-1] - A[i-2] {
            dp[i] = dp[i-1] + 1
            sum += dp[i]
        } 
    }
    return sum
}
