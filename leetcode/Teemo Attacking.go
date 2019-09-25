func findPoisonedDuration(timeSeries []int, duration int) int {
    n := len(timeSeries)
    ans := duration * n
    for i:=1; i<n; i++ {
        current := duration - (timeSeries[i] - timeSeries[i-1])
        if current > 0 {ans -= current}
    }
    return ans
}
