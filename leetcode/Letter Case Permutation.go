func letterCasePermutation(S string) []string {
	var prev = []string{""}
	for _, x := range S {
		var current []string
		X := string(x)
		for _, p := range prev {
			if unicode.IsDigit(x){
                current = append(current, p+X)
            } else if unicode.IsLower(x) {
                current = append(current, p+X, p+strings.ToUpper(X))
            } else {
                current = append(current, p+strings.ToLower(X), p+X)
            }
		}
		prev = current
	}
	return prev
}
