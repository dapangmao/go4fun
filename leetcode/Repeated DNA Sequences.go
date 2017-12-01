func findRepeatedDnaSequences(s string) []string {
    var dic = make(map[string]int)
    for i:=0; i<=len(s)-10; i++ {
        dic[s[i:i+10]]++
    }
    var res = []string{}
    for k, v := range dic {
        if v > 1 {
            res = append(res, k)
        }
    }
    return res
}
