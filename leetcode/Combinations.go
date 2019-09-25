func combine(n int, k int) [][]int {
    var res [][]int
    var dfs func(i int, path []int) 
    dfs = func(i int, path []int) {
        if len(path) == k {
            res = append(res, path)
            return
        }
        if len(path) > k {return}
        for j:=i; j<=n; j++ {
			      newpath := append([]int{j}, path...)
            dfs(j+1, newpath)
        }
    }
    dfs(1, []int{})
    return res
}
