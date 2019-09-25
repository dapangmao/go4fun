func numMatchingSubseq(S string, words []string) int {
    var res int
    var dic = make(map[string]int)
    for _, word := range words {dic[word]++}
    for word := range dic {
        if isSubsequence(word, S) {res += dic[word]}
    }
    return res
}

func isSubsequence(s string, t string) bool {
    n, m := len(s), len(t)
    if n == 0 {return true}
    var j int
    for i:=0; i<m; i++ {
        if s[j] == t[i] {
            j++
        }
        if j == n {return true}
    }
    return false
}
