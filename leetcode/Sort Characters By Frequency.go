func frequencySort(s string) string {
    var dic = make(map[byte]int)
    var bs = []byte(s)
    for _, x := range bs {
        dic[x]++
    }
    sort.Slice(bs, func(i, j int) bool {
        return float64(dic[bs[j]]) + float64(bs[j]) / 1000.0 < float64(dic[bs[i]]) + float64(bs[i]) / 1000.0 
    })
    return string(bs)
}
