func shortestCompletingWord(licensePlate string, words []string) string {
    var ldic = make(map[rune]int)
    for _, x := range strings.ToLower(licensePlate) {
        if x >= 97 && x <= 122 { 
            ldic[x]++
        }
    }
    var res string
    for _, word := range words {
        var dic = make(map[rune]int)
        for _, x := range(word) {
            dic[x]++
        }
        for k, v := range ldic {
            value, ok := dic[k]
            if !ok || value < v {goto end}
        }
        if len(res) == 0 || len(word) < len(res) {res = word}
        end:
    }
    return res
}
