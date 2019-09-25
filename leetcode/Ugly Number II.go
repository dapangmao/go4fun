func nthUglyNumber(n int) int {
    var ugly = make([]int, n)
    ugly[0] = 1
    var index2, index3, index5 int
    factor2, factor3, factor5 := 2, 3, 5
    for i:=1; i<n; i++ {
        var min = factor2
        if factor3 < min {min = factor3}
        if factor5 < min {min = factor5}
        ugly[i] = min
        if factor2 == min {
            index2++
            factor2 = 2*ugly[index2]
        }
        if factor3 == min {
            index3++
            factor3 = 3*ugly[index3]
        }
        if factor5 == min {
            index5++
            factor5 = 5*ugly[index5]
        }
    }
    return ugly[n-1]
}
