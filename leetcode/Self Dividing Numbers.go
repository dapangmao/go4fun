func selfDividingNumbers(left int, right int) []int {
    res := []int{}
    for x:=left; x<=right; x++ {
        y := x
        for y > 0 {
            current := y % 10
            if current == 0 || x % current != 0 {
                goto end
            }
            y /= 10
        }
        res = append(res, x)
        end:
    }
    return res
}
