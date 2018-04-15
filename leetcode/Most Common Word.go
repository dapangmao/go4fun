func mostCommonWord(paragraph string, banned []string) string {
    max, cand := -1, ""
    var set = make(map[string]bool)
    for _, x := range banned { set[x] = true }
    var splits = strings.Fields(paragraph)
    var dic = make(map[string]int)
    for _, x := range splits {
        var current = strings.ToLower(x)
        last := x[len(x)-1] 
        for _, y := range []byte("!?',;.") {
            if last == y {
                current = current[:len(x)-1]
                break
            }
        }
        if _, ok := set[current]; ok {continue}
        dic[current]++ 
        if dic[current] > max {
            cand = current
            max = dic[current]
        }
    }
    return cand
}
