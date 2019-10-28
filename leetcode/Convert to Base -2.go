func baseNeg2(N int) string {
    if N <= 0 {return "0"}
    if N == 1 {return "1"}
    var res []int
    rem, quo := 0, 0
    for N != 0 {
        rem = N % -2
        quo = N / -2
        if rem < 0 {
            rem += 2
            quo++
        }
        res = append(res, rem)
        N = quo
    }
	var ans strings.Builder
	for j:=len(res)-1; j>=0; j-- {
		ans.WriteString(strconv.Itoa(res[j]))
	}
	return ans.String()
}
