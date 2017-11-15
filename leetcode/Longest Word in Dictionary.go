func longestWord(words []string) string {
    set := make(map[string]int)
    for _, x := range words {
        set[x] = len(x)
    }
    ans := ""
    for _, word := range words {
        if set[word] < set[ans] || (set[word] == set[ans] && word >= ans) {
            continue
        willAdd := true
        for k:=1; k<set[word]; k++ {
            if _, ok := set[word[:k]]; !ok {
                willAdd = false
                break
            } 
        }
        if willAdd {ans = word}
    }
    return ans
}
