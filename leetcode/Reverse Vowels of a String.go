func reverseVowels(s string) string {
    var bs = []byte(s)
    i, j := 0, len(bs)-1
    for i < j {
        for i < j && !isVowel(bs[i]) {i++}
        for i < j && !isVowel(bs[j]) {j--}
        bs[i], bs[j] = bs[j], bs[i]
        i++
        j--
    }
    return string(bs)
}

func isVowel(a byte) bool {
    if a == 'a' || a == 'e' || a == 'u' || a == 'o' || a == 'i' {return true}
    if a == 'A' || a == 'E' || a == 'U' || a == 'O' || a == 'I' {return true}
    return false
}
