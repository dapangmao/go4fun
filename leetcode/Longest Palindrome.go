func longestPalindrome(s string) int {
    var set = make(map[rune]bool)
    var res = 0
    for _, x := range s {
        if _, ok := set[x]; ok {
            delete(set, x)
            res += 2
        } else {
            set[x] = true
        }
    }
    if len(set) > 0 {return res + 1}
    return res
}
