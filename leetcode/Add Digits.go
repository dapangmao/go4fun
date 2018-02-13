func addDigits(num int) int {
    for num > 9 {
        var current int
        for num > 0 { 
            current += num % 10
            num /= 10
        }
        num = current
    }
    return num
}
