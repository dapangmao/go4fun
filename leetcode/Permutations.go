func permute(nums []int) [][]int {
    n := len(nums)
    if n == 1 {return [][]int{nums}}
    var res [][]int    
    for i, x := range nums {  
        var others = append(append([]int{}, nums[:i]...), nums[i+1:]...)
        for _, y := range permute(others) {
            res = append(res, append(y, x))
        }
    }
    return res
}
