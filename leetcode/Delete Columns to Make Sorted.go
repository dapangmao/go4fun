func minDeletionSize(A []string) int {
    n, m := len(A), len(A[0])
    if n == 1 {
        return 0
    }
    var res int
    for j:=0; j<m; j++ {
        var prev byte = 'a'
        for i:=0; i<n; i++ {
            if A[i][j] < prev {
                res++
                break
            }
            prev = A[i][j]
        }
    }
    return res
} 
