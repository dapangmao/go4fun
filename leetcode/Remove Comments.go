func removeComments(source []string) []string {
	res := []string{}
	inBlock := false
	for _, line := range source {
		var current string
		right := strings.Index(line, "*/")
		double := strings.Index(line, "//")
		left := strings.Index(line, "/*")
		if !inBlock && left >= 0 && right >= 0 {
            current = line[:left] + line[right+2:]
		} else if inBlock && right >= 0 {
			current = line[right+2:len(line)]
			inBlock = false
		} else if !inBlock && double >= 0 {
			current = line[:double]
		} else if !inBlock && left >= 0 {
			current = line[:left]
			inBlock = true
		} else if !inBlock {
			current = line
		}
        if len(current) > 0 {
		    res = append(res, current)
        }
	}
	return res
}
