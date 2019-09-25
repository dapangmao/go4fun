func isPerfectSquare(num int) bool {
    var sqrt = int(math.Sqrt(float64(num)))
    return sqrt * sqrt == num
}
