func numJewelsInStones(J string, S string) int {
    var dic = make(map[rune]int)
    for _, x := range J {
        dic[x] = 0
    }
    var res int
    for _, x := range S {
        if _, ok := dic[x]; ok {
            res++
        }
    }
    return res
}
