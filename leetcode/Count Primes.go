func countPrimes(n int) int {
    var dp = make([]bool, n)
    if n <= 2 {return 0}
    var res int
    for i:=2; i<n; i++ {
        if dp[i] == true {continue}
        for j:=2; i*j<n; j++ {
            dp[i*j] = true
        }
        res++
    }
    return res
}
