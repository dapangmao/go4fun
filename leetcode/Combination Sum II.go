func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var res [][]int
    var dic = make(map[string]bool)
    var dfs func(i, sum int, path []int)
    dfs = func(i, sum int, path []int) {
        if sum == target {
            var key = ""
            for _, x := range path {key += strconv.Itoa(x) + "-"}  // since length-varialbe array cannot be map key, use string instead
            if !dic[key] {
                res = append(res, path)
                dic[key] = true
            }
            return
        }
        for j:=i+1; j<len(candidates); j++ {
            var newsum = sum + candidates[j]
            if newsum > target {break}
            newpath := append(append([]int{}, path...), candidates[j]) // have to exlicitly claim a new array
            dfs(j, newsum, newpath)
        }
    }
    dfs(-1, 0, []int{})
    return res
}
