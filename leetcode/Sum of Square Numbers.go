func judgeSquareSum(c int) bool {
    var cand []int{0}
    for i:=1; i*i<=c; i++ {
        cand = append(cand, i*i)
    }
    i, j := 0, len(cand)-1
    for i <= j {
        sum := cand[i] + cand[j]
        if sum == c {return true}
        if sum < c {
            i++
        } else {
            j--
        }
    }
    return false
}
