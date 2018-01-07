func findWords(words []string) []string {
    var dict = make(map[rune]int)
    var count = 0
    for _, c := range "qwertyuiop asdfghjkl zxcvbnm" {
        if count == ' ' {
            count++
        } else {
            dict[c] = count
        }
    }
    var res []string
    for _, w := range words {
        var current = strings.toLowerCase(w)
        var first = dict[current[0]]
        for _, c := range current {
            if first != dict[c] {goto end}
        }
        res = append(res, w)
        end:
    }
    return res
}
