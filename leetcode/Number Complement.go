func findComplement(num int) int {
    var res int
    i := 1
    for ;num > 0; num >>= 1 {
        if num & 1 == 0 {
            res += i
        }
        i *= 2
    }
    return res
}
