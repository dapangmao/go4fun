func repeatedSubstringPattern(s string) bool {
    n := len(s)
    for i:=1; i <= n/2; i++ {
        if n%i != 0 {continue}
        var fragment = []byte(s[:i])
        var dup []byte
        for j:=0; j<n/i; j++ {
            dup = append(dup, fragment...)
        }
        if string(dup) == s {return true}
    } 
    return false
}
