func countCharacters(words []string, chars string) int {
    dict := count(chars)
    res := 0
    for _, w := range words {
        current := count(w)
        for k, v := range current {
            if dict[k] < v {goto next} 
        }
        res += len(w)       
        next:
    }
    return res
}


func count(s string) map[rune]int {
    res := make(map[rune]int)
    for _, r := range s {
        res[r]++
    }
    return res
}
