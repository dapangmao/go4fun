func trailingZeroes(n int) int {
    if n <= 1 {return 0}
    var res int
    for i:=n; i>0 && i%5==0; i=i/5 {
        res++
    }
    return trailingZeroes(n-1) + res
}
