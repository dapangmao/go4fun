func titleToNumber(s string) int {
    n := len(s)
    var res int
    j := 1
    for i:= n-1; i>=0; i-- {
        var current = int(s[i] - 'A') + 1
        current *= j
        res += current
        j *= 26
    }
    return res
}
