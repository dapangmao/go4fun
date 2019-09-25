func findDisappearedNumbers(nums []int) []int {
    for _, x := range nums {
        if x < 0 {x = -x}
        if nums[x-1] > 0 {
            nums[x-1] = -nums[x-1]    
        }
    }
    var res []int
    for i, x := range nums {
        if x > 0 {res = append(res, i+1)}
    }
    return res
}
