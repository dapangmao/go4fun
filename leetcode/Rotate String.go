func rotateString(A string, B string) bool {
    if len(A) != len(B) {return false}
    newB := B + B 
    if strings.Index(newB, A) == -1 {return false}
    return true
}
