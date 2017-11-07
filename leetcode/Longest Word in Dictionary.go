func longestWord(words []string) string {
    set := make(map[string]bool)
    for _, x := range words {
        set[x] = true
    }
    ans := ""
    for _, word := range words {
        if len(word) > len(ans) || len(word) == len(ans) && word < ans {
            willAdd := true
            for k:=1; k<len(word); k++ {
                if _, ok := set[word[:k]]; !ok {
                    willAdd = false
                    break
                } 
            }
            if willAdd {ans = word}
        }
    }
    return ans
}
