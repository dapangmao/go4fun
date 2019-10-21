func tribonacci(n int) int {
    if n == 0 {return 0}
    if n < 3 {return 1}
    ppp, pp, p := 0, 1, 1
    var res int
    for i:=0; i<n-2; i++ {
        res = ppp + pp + p
        ppp, pp, p = pp, p, res
    }
    return res
}
