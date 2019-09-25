func matrixReshape(nums [][]int, r int, c int) [][]int {
    var n, m = len(nums), len(nums[0])
    if n * m != r * c {return nums}
    var res = make([][]int, r)
    for i := range res {res[i] = make([]int, c)}
    var k, l int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            res[k][l] = nums[i][j]
            l++
            if l == c {
                k++
                l = 0
            }
        }
    }
    return res
}
