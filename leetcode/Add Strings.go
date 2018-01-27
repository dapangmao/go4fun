func addStrings(num1 string, num2 string) string {
    i, j := len(num1)-1, len(num2)-1
    var sum []int
    var prev int
    for i >= 0 || j >= 0 {
        var current = prev
        if i >= 0 {
            current += int(num1[i] - '0')
            i--
        }
        if j >= 0 {
            current += int(num2[j] - '0')
            j--
        }
        sum = append(sum, current%10)
        prev = current / 10
    }
    if prev > 0 {sum = append(sum, prev)}
    var res string
    for i:=len(sum)-1; i>=0; i-- {
        res += strconv.Itoa(sum[i])
    }
    return res
}
