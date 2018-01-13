func findLengthOfLCIS(nums []int) int {
    var res, current int
    for i, x := range nums {
        if i == 0 || x > nums[i-1]{
            current++
        } else {
            if current > res {
                res = current
            }
            current = 1
        }
    }
    if current > res {return current}
    return res
}
