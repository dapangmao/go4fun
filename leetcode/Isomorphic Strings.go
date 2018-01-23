func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {return false}
    return isOnewayIsomorphic(s, t) && isOnewayIsomorphic(t, s)
}

func isOnewayIsomorphic(s, t string) bool {
    var dic = make(map[byte]byte)
    for i, x := range []byte(s){
        if val, ok := dic[x]; ok {
            if val != t[i] {return false}
        } else {
            dic[x] = t[i]
        }
    }
    return true
}
