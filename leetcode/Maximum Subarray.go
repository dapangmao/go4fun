func maxSubArray(nums []int) int {
    var res = math.MinInt32
    var current int
    for _, num := range nums {
        current += num
        if current > res {res = current}
        if current < 0 {current = 0}
    }
    return res
}
