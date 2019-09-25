func firstUniqChar(s string) int {
    m := make(map[rune]int)
    for _, x := range s {m[x]++}
    for i, x := range s {
        if m[x] == 1 {return i}
    }
    return -1
}
