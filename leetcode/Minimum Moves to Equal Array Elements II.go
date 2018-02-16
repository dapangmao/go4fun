func minMoves2(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    var median int
    if n % 2 == 0 {
        median = (nums[n/2-1] + nums[n/2]) / 2
    } else {
        median = nums[n/2]
    }
    var res int
    for _, num := range nums {
        var current = median - num
        if current < 0 {current = -current}
        res += current
    }
    return res
} 
