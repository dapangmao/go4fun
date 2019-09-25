func isPalindrome(s string) bool {
	var bytes []int
	for _, x := range []byte(s) {
		var cand1 = x - 'a'
		var cand2 = x - 'A'
		var cand3 = x - '0'
		if cand1 <= 25 {
			bytes = append(bytes, int(cand1))
		} else if cand2 <= 25 {
			bytes = append(bytes, int(cand2))
		} else if cand3 <= 9 {
			bytes = append(bytes, int(cand3))
		}
	}
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {return false}
		i++
		j--
	}
	return true
}
