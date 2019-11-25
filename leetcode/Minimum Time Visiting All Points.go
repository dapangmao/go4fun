func minTimeToVisitAllPoints(points [][]int) int {
    var res int
    for i:=1; i<len(points); i++ {
        res += getDist(points[i], points[i-1])
    }
    return res
}


func min(a, b int) int {
    if a < b {return a}
    return b
}
func abs(a, b int) int {
    res := a - b
    if res < 0 {return -res}
    return res
}


func getDist(p1, p2 []int) int {
    width, length := abs(p1[0], p2[0]), abs(p1[1], p2[1])
    return min(width, length) + abs(width, length)
}
