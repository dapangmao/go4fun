func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    if n == 0  {return false}
    m := len(matrix[0])
    if m == 0  {return false}
    lo, hi := 0, n-1
    for lo <= hi {
        mid := (hi+lo) / 2
        var idx = sort.SearchInts(matrix[mid], target)
        if target == matrix[mid][idx] {return true}
        if idx == 0 {
            hi = mid - 1
        } else if idx == m {
            lo = mid + 1
        }
    }
    return false
}
