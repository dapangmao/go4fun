func lengthOfLastWord(s string) int {
    n := len(s)
    for i:=len(s)-1; i>=0; i-- {
        if s[i] != ' ' {
            break
        }
        n -= 1 
    }
    for i:=n-1; i>=0; i-- {
        if s[i] == ' ' {
            return n-1-i   
        }
    }
    return n
}
