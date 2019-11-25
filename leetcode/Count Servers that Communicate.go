func countServers(grid [][]int) int {
    rows, cols := make(map[int]int), make(map[int]int)
    n, m := len(grid), len(grid[0])
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] == 1 {
                rows[i]++
                cols[j]++
            }
        }
    }
    var res int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] == 0 {continue}
            if rows[i] > 1 || cols[j] > 1 {res++}
        }
    }
    return res
}
