func letterCombinations(digits string) []string {
    if digits == "" {return []string{}}
    var dic = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    var res = []string{""}
    for _, d := range []byte(digits) {
        var options = dic[int(d-'0')]
        var current []string
        for _, s := range res {
            for _, o := range options {
                current = append(current, s+string(o))
            }
        }
        res = current
    }
    return res
}
