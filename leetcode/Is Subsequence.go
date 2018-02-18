func isSubsequence(s string, t string) bool {
    n, m := len(s), len(t)
    if n == 0 {return true}
    var j int
    for i:=0; i<m; i++ {
        if s[j] == t[i] {
            j++
        }
        if j == n {return true}
    }
    return false
}
