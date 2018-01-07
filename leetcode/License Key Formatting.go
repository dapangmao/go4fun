func licenseKeyFormatting(S string, K int) string {
    var res []byte
    var count = K
    for i:=len(S)-1; i>=0; i-- {
        var c = S[i]
        if c == '-' {continue}
        res = append([]byte{c}, res...)
        count--
        if i > 0 && count == 0 {
            res = append([]byte{'-'}, res...)
            count = K
        }
    }
    count = 0
    for _, c := range res {
        if c != '-' {break}
        count++
    }
    return strings.ToUpper(string(res[count:]))
}
