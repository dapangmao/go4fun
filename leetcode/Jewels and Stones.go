func numJewelsInStones(J string, S string) int {
    var set = make(map[rune]bool)
    for _, x := range J {
        set[x] = true
    }
    var res int
    for _, x := range S {
        if set[x] {res++}
    }
    return res
}
