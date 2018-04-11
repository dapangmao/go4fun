func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    if n == 0 {return false}
    m := len(matrix[0])
    if m == 0 {return false}
    i, j := 0, m-1
    for i < n && j >= 0 {
        current := matrix[i][j]
        if current == target {
            return true
        }
        if target > current {
            i++
        } else {
            j--
        }
    }
    return false
}
