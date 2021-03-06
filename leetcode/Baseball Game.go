func calPoints(ops []string) int {
    var stack []int
    for _, op := range ops {
        n := len(stack)
        if op == "+" {
            stack = append(stack, stack[n-1] + stack[n-2])
        } else if op == "C" {
            stack = stack[:n-1]
        } else if op == "D" {
            stack = append(stack, 2*stack[n-1])
        } else {
            e, _ := strconv.Atoi(op)
            stack = append(stack, e)
        }
    }
    var res = 0
    for _, x := range stack {
        res += x
    }
    return res
}
