func countSubstrings(s string) int {
    count := func(left, right int) int {
        ans := 0
        for left >= 0 && right < len(s) && s[left] == s[right] {
            ans += 1
            left -= 1
            right += 1
        }
        return ans
    }
    var res int
    for i:=0; i<len(s); i++ {
        res += count(i, i) + count(i, i+1)
    }
    return res
}
