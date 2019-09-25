func solveEquation(equation string) string {
	r, _ := regexp.Compile(`(=)|([-+]?)(\d*)(x?)`)
	x, a := 0, 0
	side := 1
	var splits = r.FindAllStringSubmatch(equation, -1)
	for _, row := range splits {
		eq, _sign, _num, isx := row[1], row[2], row[3], row[4]
		sign := 1
		if _sign == "-" {sign = -1}
		num, err2 := strconv.Atoi(_num)
		if eq == "=" {
			side = -1
		} else if isx == "x" {
			if err2 != nil {
				x += side  * sign * num
			} else {
				x += side * sign
			}
		} else if err2 == nil {
			a -= side * sign * num
		}
	}
	if a == 0 {return "Infinite solutions"}
	if x == 0 {return "No solution"}
	return strconv.Itoa(a/x)
}
