func suggestedProducts(products []string, searchWord string) [][]string {
    var res [][]string
    sort.Strings(products)
    var prev = products
    for i, b := range []byte(searchWord) {
        var current []string
        for _, w := range prev {
            if i >= len(w) {continue}
            if w[i] == b {current = append(current, w)}
        }
        if len(current) > 3 {
            res = append(res, current[:3])
        } else {
            res = append(res, current)
        }
        prev = current
    }
    return res
}
