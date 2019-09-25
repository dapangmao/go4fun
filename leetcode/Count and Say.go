func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = nextState(res)
	}
	return res
}

func nextState(s string) string {
	res, prev := "", ""
	count := 0
	for _, x := range s {
		current := string(x)
		if prev != current  {
			if count != 0 {
				res += strconv.Itoa(count) + prev
			}
			prev = current
			count = 1
		} else {
			count += 1
		}
	}
	return res + strconv.Itoa(count) + prev
}
