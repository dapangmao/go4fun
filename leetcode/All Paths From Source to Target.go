func allPathsSourceTarget(graph [][]int) [][]int {
    var res [][]int
    var n = len(graph) - 1
    var dfs func(i int, path []int)
    dfs = func(i int, path []int) {
        for _, x := range graph[i] {
            newPath := append(append([]int{}, path...), i)
            if x == n {
                newPath = append(newPath, n)
                res = append(res, newPath)
            } else {
                dfs(x, newPath)
            }
        }
    }
    dfs(0, []int{})
    return res
}
