func oddCells(n int, m int, indices [][]int) int {
    mat := make([][]int, n)
    for i := range mat {
        mat[i] = make([]int, m)
    }
    rows, cols := make(map[int]int), make(map[int]int)
    for _, indice := range indices {
        r, c := indice[0], indice[1]
        rows[r]++
        cols[c]++
    }
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            mat[i][j] += rows[i] + cols[j]
        }
    }
    var res int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if mat[i][j] % 2 == 1 {res++}
        }
    }
    return res
}
