func selfDividingNumbers(left int, right int) []int {
    res := []int{}
    for x:=left; x<=right; x++ {
        is_dividable := true
        y := x
        for y > 0 {
            current := y % 10
            if current == 0 || x % current != 0 {
                is_dividable = false
                break 
            }
            y /= 10
        }
        if is_dividable {
            res = append(res, x)
        }
    }
    return res
}
