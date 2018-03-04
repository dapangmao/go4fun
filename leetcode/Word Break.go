func wordBreak(s string, wordDict []string) bool {
    set := make(map[string]bool)
    for _, word := range wordDict {set[word] = true}
    dp := make([]bool, len(s))
    for i := range s {
        for j:=0; j<=i; j++ {
            if (j == 0 || dp[j-1]) && set[s[j:i+1]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)-1]
}
