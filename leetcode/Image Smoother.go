func imageSmoother(M [][]int) [][]int {
    n, m := len(M), len(M[0])
    res := make([][]int, n)
    for i := range M {
        res[i] = make([]int, m)
    }
    var current, count int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            current = 0
            count = 0
            for _, l := range []int{-1, 0, 1} {
                for _, k := range []int{-1, 0, 1} {
                    r, c := i+l, j+k
                    if r >=0 && r < n && c >= 0 && c <m {
                        current += M[r][c]
                        count++
                    }
                }
            }
            res[i][j] = int(math.Floor(float64(current/count)))
        }
    }
    return res
}
