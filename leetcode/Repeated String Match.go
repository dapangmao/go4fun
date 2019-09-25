func repeatedStringMatch(A string, B string) int {
    var increase = A
    var res int = 1
    for {
        if strings.Contains(increase, B) {
            return res
        } else if len(increase) > 2*len(B) {
            break
        } else {
            increase += A
            res++
        }
    }
    return -1
}
