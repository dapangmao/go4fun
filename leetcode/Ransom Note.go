func canConstruct(ransomNote string, magazine string) bool {
    var dic = make(map[rune]int)
    for _, x := range magazine {
        dic[x]++
    }
    for _, x := range ransomNote {
        dic[x]--
        if dic[x] < 0 {return false}
    }
    return true
}
