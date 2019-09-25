func largestPerimeter(A []int) int {
    sort.Ints(A)
    n := len(A)
    for i:=n-3; i>=0; i-- {
        a, b, c := A[i], A[i+1], A[i+2]
        if a + b > c {
            return a + b + c
        }
    }
    return 0
}
