func groupAnagrams(strs []string) [][]string {
    var dic = make(map[string][]string)
    for _, str := range strs {
        var key = makeKey(str)
        dic[key] = append(dic[key], str)
    }
    var res [][]string
    for _, val := range dic {
        res = append(res, val)
    }
    return res
}

func makeKey(s string) string {
    var bucket = make([]int, 26)
    for _, x := range s {
        bucket[x - 'a']++
    }
    var res string
    for i, x := range bucket {
        if x <= 0 {continue}
        res += fmt.Sprintf("%d-%d/", i, x)
    }
    return res
}
