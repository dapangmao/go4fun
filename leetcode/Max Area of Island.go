func maxAreaOfIsland(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    var res int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] != 1 {continue}
            queue := [][]int{[]int{i, j}}
            grid[i][j] = -1
            current := 0
            for len(queue) > 0 {
                var pop = queue[0]
                queue = queue[1:]
                var a, b = pop[0], pop[1]
                current++
                var options = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
                for _, option := range options {
                    var r = a + option[0]
                    var c = b + option[1]
                    if r >=0 && r < n && c >= 0 && c < m && grid[r][c] == 1 {
                        queue = append(queue, []int{r, c})
                        grid[r][c] = -1
                    }
                }
            }
            if current > res {res = current}
        }
    }
    return res
}
