func maxProduct(words []string) int {
    hash := make(map[string]int)
    var max int
    for _, word := range words {
        for k, v := range hash {
            if v >= len(word) {continue}
            var current = numIfDiff(k, word)
            if current > v {
                hash[k] = v
                if current*len(k) > max {max = current*len(k)}
            }
        }
        hash[word] = 0
    }
    return max
}

func numIfDiff(a, b string) int {
    var bucket [26]int
    for _, x := range []byte(a) {
        bucket[int(x-'a')]++
    }
    for _, x := range []byte(b) {
        if bucket[int(x-'a')] > 0 {return -1}
    }
    return len(b)
}
