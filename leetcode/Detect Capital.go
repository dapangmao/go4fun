func detectCapitalUse(word string) bool {
    var allLower, allUpper, firstUpper = 0, 0, 0
    var n = len(word)
    for i, c := range word {
        if c >= 'A' && c <= 'Z' {
            allUpper++
            if i == 0 {firstUpper = 1}
        } else {
            allLower++
        }
    } 
    if allLower == n || allUpper == n || (firstUpper == 1 && allLower == n-1) {return true}
    return false
}
