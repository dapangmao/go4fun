func isValid(s string) bool {
    var pairs = map[rune]rune{')': '(', ']': '[', '}': '{' }
    var stack []rune
    var popout rune
    for _, x := range s {
        if x == '(' || x == '[' || x == '{' {
            stack = append(stack, x)
        } else {
            var n = len(stack)
            if n == 0 {return false}
            popout, stack = stack[n-1], stack[:n-1]
            if popout != pairs[x] {return false}
        }
    }
    return len(stack) == 0
}
