func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var res [][]int
    var dfs func(i, sum int, path []int)
    dfs = func(i, sum int, path []int) {
        if sum == target {
            res = append(res, path)
            return
        }
        for j:=i; j<len(candidates); j++ {
            var newsum = sum + candidates[j]
            if newsum > target {break}
            newpath := append([]int{candidates[j]}, path...) // have to exlicitly claim a new array
            dfs(j, newsum, newpath)
        }
    }
    dfs(0, 0, []int{})
    return res
}
