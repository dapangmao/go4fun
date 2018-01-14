func getRow(rowIndex int) []int {
    var res = []int{1}
    for i:=1; i<=rowIndex; i++ {
        res = append(res, 1)
        var prev = 1
        for j:=1; j<len(res)-1; j++ {
            current := res[j]
            res[j] += prev
            prev = current
        }
    }
    return res
}
