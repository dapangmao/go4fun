func firstUniqChar(s string) int {
    m := make(map[rune]int)
    for _, x := range s {
        if val, ok := m[x]; ok {
            m[x] = val + 1
        } else {
            m[x] = 1
        }
    }
    for i, x := range s {
        if m[x] == 1 {return i}
    }
    return -1
}
