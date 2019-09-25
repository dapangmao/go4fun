func multiply(num1 string, num2 string) string {
    var len1, len2 = len(num1), len(num2)
    var product = make([]int, len1+len2)
    for i := len1 - 1; i >= 0; i-- {
        for  j := len2 - 1; j >= 0; j-- {
            var index = len1 + len2 - i - j - 2
            product[index] += int(num1[i] - '0') * int(num2[j] - '0')
            product[index + 1] += product[index] / 10;
            product[index] %= 10;
        }
    }
    var stringBuilder []byte
    for i := len(product) - 1; i > 0; i-- {
        if len(stringBuilder) == 0 && product[i] == 0 {
            continue
        }
        stringBuilder = append(stringBuilder, byte(product[i]))
    }
    stringBuilder = append(stringBuilder, byte(product[0]))
    var res string
    for _, x := range stringBuilder {res += fmt.Sprint(x)}
    return res
}
