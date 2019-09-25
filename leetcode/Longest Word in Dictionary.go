func longestWord(words []string) string {
    set := make(map[string]bool)
    for _, x := range words {set[x] = true}
    ans := ""
    for _, word := range words {
        if len(word) > len(ans) || len(word) == len(ans) && word < ans {
            for k:=1; k<len(word); k++ {
                if !set[word[:k]] {goto end} 
            }
            ans = word
            end:
        }
    }
    return ans
}
