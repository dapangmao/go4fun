func longestPalindrome(s string) string {
    var res string
    var n = len(s)
    var getPd func(i, j int)
    getPd = func(i, j int) {
        for {
            if i<0 || j>=n || s[i] != s[j] {break}
            if j-i+1 > len(res) {res = s[i:j+1]}
            i--
            j++
        }
    }
    for k := range s {
        getPd(k, k)
        getPd(k, k+1)
    }
    return res
}
