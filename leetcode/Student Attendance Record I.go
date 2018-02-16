func checkRecord(s string) bool {
    n := len(s)
    hasA := false
    for i:=0; i<n; i++ {
        if s[i] == 'A' {
            if hasA {
                return false
            } else {
                hasA = true
            }
        }
        if i > 0 && i < n-1 && s[i] == 'L' && s[i-1] == 'L' && s[i+1] == 'L' {
            return false
        }
    }
    return true
}
