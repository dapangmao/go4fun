func ambiguousCoordinates(S string) []string {
    var res []string
    n := len(S)
    for i:=2; i<n-1; i++ {
        for _, left := range getOpts(S, 1, i) {
            for _, right := range getOpts(S, i, n-1) {
                res = append(res, fmt.Sprintf("(%s, %s)", left, right))
            }
        }
    }
    return res
}


func getOpts(S string, i, j int) []string {
	var res []string
	for d:=1; d <=j-i; d++ {
		left, right := S[i:i+d], S[i+d:j]
		if len(right) > 0 && right[len(right)-1] == '0' {continue}
		if left != "0" && left[0] == '0' {continue}
		var current = left + right
		if d < j-i {
			current = left + "." + right
		}
		res = append(res, current)
	}
	return res
}
