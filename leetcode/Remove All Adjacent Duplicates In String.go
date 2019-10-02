func removeDuplicates(S string) string {
    var current = []byte(S)
    for {
        lastLen := len(current)
        current = remove(current)
        if lastLen == len(current) { break }
    }
    return string(current)
}

func remove(text []byte) []byte {
    var buf []byte
    for i, b := range text {
        if (i > 0 && text[i-1] == b) || (i<len(text)-1 && text[i+1] == b) {continue}
        buf = append(buf, b)
    }
    return buf
} 
