func hasAlternatingBits(n int) bool {
    prev := n & 1
    n >>= 1
    for n > 0 {
        var current = n & 1
        if current == prev {return false}
        prev = current
        n >>= 1
    }
    return true
}
