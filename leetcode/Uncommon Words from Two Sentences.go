func uncommonFromSentences(A string, B string) []string {
    var res []string
    findDiff := func(a, b string) {
        hash := make(map[string]int)
        for _, str := range strings.Fields(a) {
            hash[str]++
        }
        for _, str := range strings.Fields(b) {
            delete(hash, str)
        } 
        for k, v := range hash {
            if v == 1 {
                res = append(res, k)
            }
        }
    }
    findDiff(A, B)
    findDiff(B, A)
    return res
}
