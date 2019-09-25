func convertToTitle(n int) string {
    var res []byte
    var current int
    for n > 0 {
        n -= 1
        current = 'A' + n % 26
        res = append([]byte{byte(current)}, res...)
        n /= 26
    }
    return string(res)
}
