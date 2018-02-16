func findAnagrams(s string, p string) []int {
    hash := make([]int, 26)
    var res []int
    n := len(p)
    if n > len(s) {return res}
    for _, x := range []byte(p) {
        hash[int(x-'a')]++
    }
    for _, x := range []byte(s[:n-1]) {
        hash[int(x-'a')]--
    }
    for i:=0; i<len(s)-n+1; i++{
        var inchar = s[i+n-1]
        hash[int(inchar-'a')]--
        if checkHash(hash) {res = append(res, i)}
        var outchar = s[i]
        hash[int(outchar-'a')]++
    }
    return res
}

func checkHash(hash []int) bool {
    for _, x := range hash {
        if x != 0 {return false}
    }
    return true
}
