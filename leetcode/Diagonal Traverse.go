func findDiagonalOrder(matrix [][]int) []int {
    n, m := len(matrix), len(matrix[0])
    var dirs [][]int
    for j:=0; j<m; j++ {
        dirs = append(dirs, []int{0, j})
    }
    for i:=1; i<n; i++ {
        dirs = append(dirs, []int{i, m-1})
    }
    var res []int
    for k, dir := range dirs {
        i, j := dir[0], dir[1]
        var current []int
        for i < n && j >=0  {
            current = append(current, matrix[i][j])
            i++
            j--
        }
        if k%2 == 0 {reverse(current[:]))}
        res = append(res, current...)
    }
    return res
}
