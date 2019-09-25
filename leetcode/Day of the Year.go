func isLeap(year int) bool {
    return (year % 4 == 0 && year % 100 != 0 ) || year % 400 == 0
}


func p(s string) int {
    res, _ := strconv.Atoi(s)
    return res
}

var days = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func init() {
    var sum int
    for i, x := range days {
        sum += x
        days[i] = sum
    }
}

func dayOfYear(date string) int {
    dateArr := strings.Split(date, "-")
    year, month, day := p(dateArr[0]), p(dateArr[1]), p(dateArr[2])
    res := days[month-1] + day
    if isLeap(year) && month > 2 {res++}
    return res
}
