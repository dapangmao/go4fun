func validPalindrome(s string) bool {    
    isPal := func(s []byte) bool {
        for i:=0; i<len(s)/2; i++ {
            if s[i] != s[len(s)-1-i] {return false}
        }
        return true
    }
    for i,j := 0,len(s)-1; i<j; i,j=i+1,j-1 {
        if s[i] != s[j] {
            op1 := append([]byte(s[:i]), s[i+1:]...)
            op2 := append([]byte(s[:j]), s[j+1:]...)
            return isPal(op1) || isPal(op2)
        }
    }
    return true
}
