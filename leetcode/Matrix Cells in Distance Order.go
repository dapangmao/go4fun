func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    var res = make([][]int, R*C)
    
    dist := func(slice []int) int {
        x, y := slice[0] - r0, slice[1] - c0
        if x < 0 {x = -x}
        if y < 0 {y = -y}
        return x + y
    }
    k := 0
    for i:=0; i<R; i++ {
        for j:=0; j<C; j++ {
            res[k] = []int{i, j}
            k++
        }
    }
    sort.Slice(res, func(i, j int) bool {
        return dist(res[i]) < dist(res[j])
    })
    return res
}
