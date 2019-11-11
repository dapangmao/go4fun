
func closedIsland(grid [][]int) int {
    var res int
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    n, m := len(grid), len(grid[0])
    
    var dfs func(x, y int) bool
    
    dfs = func(x, y int) bool {
        if x < 0 || y < 0 || x >= n || y >= m {return false}
        if grid[x][y] == 1 {return true}
        grid[x][y] = 1
        res := true
        for _, dir := range dirs {
            res = res && dfs(x+dir[0], y+dir[1]) 
        }
        return res
    }
    
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] == 0 {
                if dfs(i, j) {res++}
            }
        }
    }
    return res
}
