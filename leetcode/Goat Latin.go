var vowels = map[byte]bool{'a': true, 'A': true, 'e': true, 'E': true, 'o': true, 'O': true, 'i': true, 'I': true, 'u': true, 'U': true}

func toGoatLatin(S string) string {
    arr := strings.Fields(S)
    parse := func(w string) string {
        first, rest := w[0], w[1:]
        if vowels[first] {
            return w + "ma"
        }
        return rest + string(first) + "ma"
    }
    var res []string
    var appendix = "a"
    for _, x := range arr {
        res = append(res, parse(x) + appendix)
        appendix += "a"
    }
    return strings.Join(res, " ")
}
