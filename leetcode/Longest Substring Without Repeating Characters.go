func lengthOfLongestSubstring(s string) int {
    bucket := make([]int, 256)
    j := 0
    var max int
    
    ok := func() bool {
        for _, x := range bucket {
            if x > 1 {return false}
        }
        return true
    }
    
    for i, x := range []byte(s) {
        bucket[int(x-'a')]++
        for j < i && !ok() {
            pop := s[j]
            bucket[int(pop-'a')]--
            j++
        }
        if i - j + 1 > max {max = i-j+1}
    }
    return max
}
