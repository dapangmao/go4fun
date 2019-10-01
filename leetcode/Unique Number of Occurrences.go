func uniqueOccurrences(arr []int) bool {
    dict := make(map[int]int)
    for _, x := range arr {
        dict[x]++
    }
    count := make(map[int]bool)
    for _, v := range dict {
        count[v] = true
    }
    return len(dict) == len(count)
}
