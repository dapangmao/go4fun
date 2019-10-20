func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool {
		return strings.Count(folder[i], "/") < strings.Count(folder[j], "/")
	})
	dict := make(map[string]bool)
	for _, s := range folder {
		current := getPaths(s)
		for _, x := range current {
			if dict[x] {goto next}
		}
		dict[current[len(current)-1]] = true
	next:
	}
	var ans []string
	for k := range dict {
		ans = append(ans, k)
	}
	return ans
}

func getPaths(s string) []string {
	res := []string{s}
	for i:=len(s)-2; i>=0; i-- {
		if s[i] == '/' {
			res = append([]string{s[:i]}, res...)
		}
	}
	return res
}
