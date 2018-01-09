func countSegments(s string) int {
    var res = 0
    var prev = ' '
    for _, x := range s + " " {
        if x == ' ' && prev != ' ' {
            res++
        }
        prev = x
    }
    return res
}
