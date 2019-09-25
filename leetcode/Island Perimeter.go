func islandPerimeter(grid [][]int) int {
    var n, m = len(grid), len(grid[0])
    var res int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] != 1 {continue}
            if i+1 >= n || grid[i+1][j] == 0 {res++}
            if i-1 < 0 || grid[i-1][j] == 0 {res++}
            if j+1 >= m || grid[i][j+1] == 0 {res++}
            if j-1 < 0 || grid[i][j-1] == 0 {res++}
        }
    }
    return res
}
