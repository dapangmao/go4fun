func isAnagram(s string, t string) bool {
    var m = make(map[rune]int)
    for _, x := range s {
        m[x]++
    }
    for _, x := range t {
        m[x]--
    }
    for _, v := range m {
        if v != 0 {return false}
    }
    return true
}
