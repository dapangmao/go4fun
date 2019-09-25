func minFallingPathSum(A [][]int) int {
    n := len(A)
    for i:=1; i<n; i++ {
        for j:=0; j<n; j++ {
            cand := A[i-1][j]
            if j > 0 && A[i-1][j-1] < cand {cand = A[i-1][j-1]}
            if j < n-1 && A[i-1][j+1] < cand {cand = A[i-1][j+1]}
            A[i][j] += cand 
        }
    }
    res := 34243242
    for i:=0; i<n; i++ {
        if A[n-1][i] < res {res = A[n-1][i]}
    }
    return res
}
