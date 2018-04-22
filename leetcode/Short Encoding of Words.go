func minimumLengthEncoding(words []string) int {
    set := make(map[string]bool)
    for _, word := range words {set[word] = true}
    for _, word := range words {
        for i:=1; i<len(word); i++ {
            delete(set, word[i:])
        }
    }
    var res = len(set) 
    for k := range set {res += len(k)}
    return res
}
