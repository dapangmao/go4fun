func findDuplicates(nums []int) []int {
    var res []int
    for _, num := range nums {
        if num < 0 {num = -num}
        if nums[num-1] < 0 {
            res = append(res, num)
        } else {
            nums[num-1] = -nums[num-1]
        }
    }
    return res
}
