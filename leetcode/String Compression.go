func compress(chars []byte) int {
	var strs []string
	var n = len(chars)
	if n == 0 {return 0}
	var prev = chars[0]
	var count int
	for i:=0; i<=n; i++ {
		if i == n || chars[i] != prev {
			if count == 1 {
				strs = append(strs, string(prev))
			} else {
				strs = append(strs, string(prev), strconv.Itoa(count))
			}
			if i < n {prev = chars[i]}
			count = 1
		} else {
			count++
		}
	}
	var res int
	for _, s := range strs {
		res += len(s)
	}
	return res
}
