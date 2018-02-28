func convertToTitle(n int) string {
    var res []byte
    for n > 0 {
        n -= 1
        current := byte('A' + n % 26)  // implicit transformation
        res = append([]byte{current}, res...)
        n /= 26
    }
    return string(res)
}
