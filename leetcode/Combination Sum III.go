var K, N int
var res [][]int

func combinationSum3(k int, n int) [][]int {
    N = n
    K = k
    dfs(0, 0, []int{})
    return res
}

func dfs(i, sum int, path []int) {
    if len(path) == K {
        if sum == N {res = append(res, path)}
        return
    }
    for j:=i+1; j<=9; j++ {
        if j + sum <= N {
            var current = append(path, j)
            dfs(j, sum+j, current)
        }
    }
}
