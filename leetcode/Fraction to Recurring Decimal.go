func fractionToDecimal(numerator int, denominator int) string {
	var res string
	if numerator * denominator < 0 {res += "-"}
	if numerator < 0 {numerator = - numerator}
	if denominator < 0 {denominator = -denominator}
	res += strconv.Itoa(numerator / denominator)
	var current = numerator % denominator
	if current == 0 {return res}
	res += "."
	dic := make(map[int]int)
	for {
		if val, ok := dic[current]; ok {
			return fmt.Sprintf("%s(%s)", res[:val], res[val:])
		}
		dic[current] = len(res)
		current *= 10
		res += strconv.Itoa(current / denominator )
		current %= denominator;
		if current == 0 {return res}
	}
	return ""
}
