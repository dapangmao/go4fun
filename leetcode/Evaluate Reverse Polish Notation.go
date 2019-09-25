func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        t, err := strconv.Atoi(token)
        if len(stack) == 0 || err == nil {
            stack = append(stack, t)
            continue
        } 
        n := len(stack)
        e1, e2 := stack[n-2], stack[n-1]
        stack = stack[:n-2]
        var current int
        if token == "+" {
            current = e1 + e2
        } else if token == "-" {
            current = e1 - e2
        } else if token == "*" {
            current = e1 * e2 
        } else {
            current = e1 / e2
        }
        stack = append(stack, current)
    }
    return stack[0]
}
