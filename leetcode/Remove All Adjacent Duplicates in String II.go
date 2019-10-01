func removeDuplicates(s string, k int) string {
	if len(s) == 0 {return s}
	var current = []byte(s)
	for {
		lastLen := len(current)
		current = dfs(current, k)
		if len(current) == lastLen {break}
	}
	return string(current)
}


func dfs(t []byte, k int) (buf []byte) {
	countPrev, bytePrev := 1, t[0]
	for i:=1; i<=len(t); i++ {
		if i == len(t) || t[i] != bytePrev {
			countPrev = countPrev % k
			if countPrev < k {
				for j:=0; j<countPrev; j++ {buf = append(buf, bytePrev)}
			}
			countPrev = 1
		} else {
			countPrev++
		}
		if i < len(t) {bytePrev = t[i]}
	}
	return
}
