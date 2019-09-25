func minimumTotal(triangle [][]int) int {
    var n = len(triangle)
    for i:=1; i<n; i++ {
        for j := range triangle[i] {
            var current = 3243242343243
            if j-1>=0 && triangle[i-1][j-1] < current {current = triangle[i-1][j-1]}
            if j < i && triangle[i-1][j] < current {current = triangle[i-1][j]}
            triangle[i][j] += current
        }
    }
    var res = 2423423432
    for _, x := range triangle[n-1] {
        if x < res {res = x}
    }
    return res
}
