func findTheDifference(s string, t string) byte {
    var m = make([]int, 26)
    for _, x := range []byte(s) {
        m[x - 'a']++
    }
    for _, x := range []byte(t) {
        m[x - 'a']--
        if m[x - 'a'] == -1 {return x}
    }
    return 0
}
