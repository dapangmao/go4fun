func numIslands(grid [][]byte) int {
    var res int
    n, m := len(grid), len(grid[0])
    if n == 0 || m == 0 {return 0}
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] != '1' {continue}
            var queue = [][]int{[]int{i, j}}
            for len(queue) > 0 {
                r, c := queue[0][0], queue[0][1]
                queue = queue[1:]
                if r+1 < n && grid[r+1][c] == '1' {
                    queue = append(queue, []int{r+1, c})
                }
                if r-1 >= 0 && grid[r-1][c] == '1' {
                    queue = append(queue, []int{r-1, c})
                }
                if c-1 >= 0 && grid[r][c-1] == '1' {
                    queue = append(queue, []int{r, c-1})
                }
                if c+1 < m && grid[r][c+1] == '1' {
                    queue = append(queue, []int{r, c+1})
                }
                grid[r][c] = '2'
            }
            res++
        }
    }
    return res
}
