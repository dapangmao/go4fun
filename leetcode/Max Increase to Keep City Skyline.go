func maxIncreaseKeepingSkyline(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    var res int
    rowMax, colMax := make([]int, n), make([]int, m)
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] > rowMax[i] {rowMax[i] = grid[i][j]}
            if grid[i][j] > colMax[j] {colMax[j] = grid[i][j]}
        }
    }
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            var current = rowMax[i]
            if current > colMax[j] {current = colMax[j]}
            if current > grid[i][j] {
                res += current - grid[i][j]
            }
        }
    }
    return res
}
