func reverseString(s string) string {
    res := []byte(s)
    i, j := 0, len(res)-1
    for i < j {
        res[i], res[j] = res[j], res[i]
        i++
        j--
    }
    return string(res)
}
