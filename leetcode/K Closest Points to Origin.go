func kClosest(points [][]int, K int) [][]int {
    sort.Slice(points, func(i, j int) bool {
        p1, p2 := points[i], points[j]
        return p1[0]*p1[0] + p1[1]*p1[1] < p2[0]*p2[0] + p2[1]*p2[1]
    })
    return points[:K]
}
