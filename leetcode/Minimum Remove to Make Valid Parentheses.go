func minRemoveToMakeValid(s string) string {
    var stack []int
    for i, x := range s {
        j := i+1
        n := len(stack)
        if x == '(' {
            stack = append(stack, j)
        } else if x == ')' {
            if n == 0 || stack[n-1] < 0 {
                stack = append(stack, -j)
            } else if stack[n-1] > 0 {
                stack = stack[:n-1]
            }
        }
    }
    var sb strings.Builder
    k := 0
    for i, b := range s {
        if k < len(stack) && i == T(stack[k]) {
            k++
        } else { 
            sb.WriteRune(b)
        }       
    }
    return sb.String()
}

func T(i int) int {
    if i < 0 {i = -i}
    return i-1
}
