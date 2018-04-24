func subsets(nums []int) [][]int {
    n := len(nums)
    var res = [][]int{}
    var dfs func(current []int, start int) 
    dfs = func(current []int, start int) {
        res = append(res, current)
        for i:=start; i<n; i++ {
            next := append([]int{nums[i]}, current...)
            dfs(next, i+1)
        }        
    }
    dfs([]int{}, 0)
    return res
}
