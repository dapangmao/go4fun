func findMaxAverage(nums []int, k int) float64 {
    var sum int
    var max int = -12312312
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
        if i >= k {
            sum -= nums[i-k]
        }
        if i >= k-1 && sum > max {max = sum}
    }
    return float64(max) / float64(k)
}
