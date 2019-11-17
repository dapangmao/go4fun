func shiftGrid(grid [][]int, k int) [][]int {
    n, m := len(grid), len(grid[0])
    prev := make([][]int, n)
    for i:=0;i<n;i++ {prev[i] = make([]int, m)}
    for z:=0; z<k; z++ {
        prev[0][0] = grid[n - 1][m - 1]
        for i:=0;i<n;i++ {
            if i < n-1 { prev[i+1][0] = grid[i][m-1]} 
            for j:=0; j<m-1; j++ {
                prev[i][j+1] = grid[i][j]
            }
        }
        prev, grid = grid, prev
        
    }
    return grid
}
