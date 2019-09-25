func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {return 0}
    sort.Slice(points, func(i, j int) bool{
        return points[i][1] < points[j][1]
    })
    current, res := points[0][1], 1
    for _, p := range points[1:] {
        if current >= p[0] {continue}
        res++
        current = p[1]
    }
    return res
}
