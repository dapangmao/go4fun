func reverseParentheses(s string) string {
	var stack []byte
	for _, b := range []byte(s) {
		if b != ')' {
			stack = append(stack, b)
		} else {
			var current []byte
			for {
				n := len(stack)
				pop := stack[n-1]
				stack = stack[:n-1]
        if pop == '(' {break}
				current = append(current, pop)
			}
			stack = append(stack, current...)
		}
	}
	return string(stack)
}
