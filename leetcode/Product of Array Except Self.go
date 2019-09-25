func productExceptSelf(nums []int) []int {
    total, n, j := 1, len(nums), 0
    var zeroCount int
    for i, num := range nums {
        if num == 0 {
            zeroCount++
            j = i
        } else {
            total *= num
        }
    }
    res := make([]int, n)
    if zeroCount > 1 {
        return res
    }
    if zeroCount == 1 {
        res[j] = total
        return res
    }
    for i, x := range nums {
        res[i] = total / x
    }
    return res
}
