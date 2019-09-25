func hammingDistance(x int, y int) int {
    var a = x ^ y
    var res int
    for a > 0 {
        if a % 2 == 1 {res++}
        a >>= 1
    }
    return res
}
