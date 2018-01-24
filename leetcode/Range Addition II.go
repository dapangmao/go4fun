func maxCount(m int, n int, ops [][]int) int {
    var x, y = m, n
    for _, op := range ops {
        if op[0] < x {x = op[0]}
        if op[1] < y {y = op[1]}
    }
    return x * y
}
