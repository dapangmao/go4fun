func countPrimeSetBits(L int, R int) int {
    var res int
    for i:=L; i<=R; i++ {
        var x = countOnes(i)
        if x == 2 || x == 3 || x == 5 || x == 7 || x == 11 || x == 13 || x == 17 || x == 19  {res++}
    }
    return res
}

func countOnes(n int) int {
    var count int
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}
