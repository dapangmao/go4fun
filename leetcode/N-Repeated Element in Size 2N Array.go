func repeatedNTimes(A []int) int {
    for i:=0; i<len(A); i++ {
        for k:=1; k<=3; k++ {
            if k+i >= len(A) {continue}
            if A[i] == A[i+k] {return A[i]}
        }
    }
    return -1
}
