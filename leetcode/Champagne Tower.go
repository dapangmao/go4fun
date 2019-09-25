func champagneTower(poured int, query_row int, query_glass int) float64 {
    var dp = make([][]float64, 102)
    for i:= range dp {dp[i] = make([]float64, i+1)}
    dp[0][0] = float64(poured)
    for r:= 0; r <= query_row; r++ {
        for c:=0; c<=r; c++ {
            q := dp[r][c] - 1.0
            if q > 0 {
                dp[r+1][c] += q / 2.0
                dp[r+1][c+1] += q / 2.0
            }
        }
    }
    res := dp[query_row][query_glass]
    if res > 1.0 {res = 1.0}
    return res
}
