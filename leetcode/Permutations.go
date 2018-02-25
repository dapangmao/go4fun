func permute(nums []int) [][]int {
    n := len(nums)
    if n == 1 {return [][]int{nums}}
    if n == 2 {
        return [][]int{nums, []int{nums[1], nums[0]}}
    }

    var res [][]int    
    for i, x := range nums {
        var others []int
        others = append(append(others, nums[:i]...), nums[i+1:]...)
        for _, y := range permute(others) {
            res = append(res, append(y, x))
        }
    }
    return res
}
