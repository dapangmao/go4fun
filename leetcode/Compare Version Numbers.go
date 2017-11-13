func compareVersion(version1 string, version2 string) int {
	v1split := strings.Split(version1, ".")
	v2split := strings.Split(version2, ".")
	n, m := len(v1split), len(v2split)
	max := n
	if m > n {max = m}
	for i:=0; i<max; i++ {
		v1current, v2current := 0, 0
		if i < n {
			v1current, _ = strconv.Atoi(v1split[i])
		}
		if i < m {
			v2current, _ = strconv.Atoi(v2split[i])
		}
		if v1current > v2current {
			return 1
		}
		if v1current < v2current {
			return -1
		}
	}
	return 0
}
