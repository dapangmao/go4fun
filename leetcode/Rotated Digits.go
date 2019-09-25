func rotatedDigits(N int) int {
    var res int
    for i:=1; i<=N; i++ {
        d := make([]int, 10)
        j := i // the scope puzzle
        for j > 0 {
            d[j%10] = 1
            j /= 10
        }
        if d[4] + d[7] + d[3] > 0 {continue}
        if d[2] + d[5] + d[6] + d[9] > 0 {
            res++
        }
    }
    return res
}
