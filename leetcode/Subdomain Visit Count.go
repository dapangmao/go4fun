func subdomainVisits(cpdomains []string) []string {
    hash := make(map[string]int)
    for _, cpdomain := range cpdomains {
        fields := strings.Fields(cpdomain)
        freq, _ := strconv.Atoi(fields[0])
        splits := strings.Split(fields[1], ".")
        for i:=0; i<len(splits); i++ {
            current := strings.Join(splits[i:], ".")
            hash[current] += freq // a default zero value
        }
    }
    var res []string
    for k, v := range hash {
        res = append(res, fmt.Sprintf("%d %s", v, k))
    }
    return res
}
// the strings package & the string function/type
// how to sort the hash map?
