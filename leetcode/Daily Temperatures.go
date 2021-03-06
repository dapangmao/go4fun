func dailyTemperatures(temperatures []int) []int {
	var n = len(temperatures)
	var stack []int
	res := make([]int, n)
	for i:=n-1; i>=0; i-- {
	for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
	    stack = stack[:len(stack)-1]
	}
	if len(stack) > 0 {
	    res[i] = stack[len(stack)-1] - i
	}
	stack = append(stack, i)
	}
	return res
}
