func accountsMerge(accounts [][]string) [][]string {
	em_to_name := make(map[string]string)
	graph := make(map[string]map[string]bool)
	for _, acc := range accounts {
		name := acc[0]
		for _, email := range acc[1:] {
      // how to simplify the 4 lines below
			if _, ok := graph[email]; ok {
				graph[email][acc[1]] = true
			} else {
				graph[email] = map[string]bool{acc[1]: true}
			}
			graph[acc[1]][email] = true
			em_to_name[email] = name
		}
	}
	seen := make(map[string]bool)
	var ans [][]string
	for email := range graph {
		if !seen[email] {
			seen[email] = true
			stack := []string{email}
			var component []string
			for len(stack) > 0 {
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				component = append(component, node)
				for nei := range graph[node] {
					if !seen[nei] {
						seen[nei] = true
						stack = append(stack, nei)
					}
				}
			}
			sort.Strings(component)
			component = append([]string{em_to_name[email]}, component...)
			ans = append(ans, component)
		}
	}
	return ans
}
