func restoreIpAddresses(s string) []string {
	var res []string

	var dfs func(a string, path string, numDot int)
	dfs = func(a string, path string, numDot int)  {
		if numDot == 3 && len(a) == 0 {
			res = append(res, path)
			return
		}
		if numDot >= 3 {return}
		for i:=1; i<4 && i<=len(a); i++ {
			cut := a[:i]
            if len(cut) > 1 && cut[0] == '0'  {continue}
			trans, _ := strconv.Atoi(cut)
			if trans > 255  {continue}
			var nextPath string
			if len(path) == 0 {
				nextPath = cut
				dfs(a[i:], nextPath, numDot)
			} else {
				nextPath = path + "." + cut
				dfs(a[i:], nextPath, numDot+1)
			}

		}
	}
	dfs(s, "", 0)
	return res
}
