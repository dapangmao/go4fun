func maxRotateFunction(A []int) int {
    n := len(A)
    max, sum := 0, 0 
    for i, x := range A {
        max += i*x 
        sum += x
    }
    var current = max 
    for i:=n-1; i>0; i-- {
        current += sum - A[i]*n
        if current > max {max = current}

    }
    return max
}
