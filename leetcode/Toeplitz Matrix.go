func isToeplitzMatrix(matrix [][]int) bool {
    n, m := len(matrix), len(matrix[0])
    for i, j:= 0, 0; i<n; i++ {
        if !check(i,j,n,m, matrix) {return false}
    }
    for i, j:= 0, 1; j<m; j++ {
        if !check(i,j,n,m, matrix) {return false}
    }
    return true
}

func check(i, j, n, m int, matrix [][]int) bool {
    var dic = make(map[int]int)
    for i < n && j < m {
        dic[matrix[i][j]]++
        if len(dic) > 1 {return false}
        i++ 
        j++
    }
    return true
}
