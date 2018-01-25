func findMaxConsecutiveOnes(nums []int) int {
    var max = -1
    var current int
    for _, x := range nums {
        if x == 0 {
            if current > max {max = current}
            current = 0
        } else {
            current++
        }
    }
    if current > max {return current}
    return max
}
