func countSquares(matrix [][]int) int {
    res := 0
    n, m := len(matrix), len(matrix[0])
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if matrix[i][j] != 1 {goto next}
            for k:=1; i+k <= n && j+k <= m; k++ {
                for x:=i; x<i+k; x++ {
                    for y:=j; y<j+k; y++ {
                        if matrix[x][y] != 1 {
                            goto next
                        }
                    }
                }
                res++
            }
            next:
        }
    }
    return res  
}
