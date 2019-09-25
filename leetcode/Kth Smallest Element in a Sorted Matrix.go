func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    var bf = make([]int, n*n) // don't use []int{} here
    var l int
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            bf[l] = matrix[i][j]
            l++
        }
    }
    sort.Slice(bf, func(i, j int) bool {
        return bf[i] < bf[j]
    })
    return bf[k-1]
}
