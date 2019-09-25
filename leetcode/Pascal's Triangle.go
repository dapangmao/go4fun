func generate(numRows int) [][]int {
    var res [][]int
    var prev = []int{1}
    if numRows > 0 {res = append(res, prev)}
    for i:=1; i<numRows; i++ {
        var n = len(prev)
        var current = make([]int, n+1)
        current[0], current[n] = 1, 1
        for j:=1; j<n; j++ {
            current[j] = prev[j-1] + prev[j]
        }
        res = append(res, current)
        prev = current
    }
    return res
}
