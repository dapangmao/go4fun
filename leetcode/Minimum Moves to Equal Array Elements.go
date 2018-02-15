func minMoves(nums []int) int {
    var min = 2 << 61
    var res int
    for _, num := range nums {
        if num < min {min = num}
    }
    for _, num := range nums {
        res += num - min
    }
    return res
}
