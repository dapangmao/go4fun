func replaceWords(dict []string, sentence string) string {
    sort.Slice(dict, func(i, j int) bool {
        return len(dict[i]) < len(dict[j])
    })
    var res []string
    var strs = strings.Split(sentence, " ")
    for _, str := range strs {
        var current = str
        for _, prefix := range dict {
            if strings.HasPrefix(str, prefix) {
                current = prefix
                break
            }
        }
        res = append(res, current)
    }
    return strings.Join(res, " ")
}
