func wordPattern(pattern string, str string) bool {
    var words = strings.Split(str, " ")
    if len(words) != len(pattern) {return false}
    var m1, m2 = make(map[rune]string), make(map[string]rune)
    for i, x := range pattern {
        word := words[i]
        if w, ok := m1[x]; ok{
            if w != word {return false}
        } else {
            m1[x] = word
        }
        if r, ok := m2[word]; ok {
            if r != x {return false}
        } else {
            m2[word] = x
        }
    }
    return true
}
