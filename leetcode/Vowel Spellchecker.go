func spellchecker(wordlist []string, queries []string) []string {
	original := make(map[string]bool)
	capitals := make(map[string]string)
	vowels := make(map[string]string)

	for _, word := range wordlist {
		original[word] = true
		capKey := strings.ToUpper(word)
		if _, ok := capitals[capKey]; !ok {
			capitals[capKey] = word
		}
		vowKey := makeVowKey(word)
		if _, ok := vowels[vowKey]; !ok {
			vowels[vowKey] = word
		}
	}
	var res []string
	for _, q := range queries {
		var current string
		if original[q] {
			current = q
		} else if w, ok := capitals[strings.ToUpper(q)]; ok {
			current = w
		} else if w, ok := vowels[makeVowKey(q)]; ok {
			current = w
		}
		res = append(res, current)
	}
	return res
}

func makeVowKey(s string) string {
	var sb strings.Builder
    for i, b := range bytes.ToLower([]byte(s)) {
		if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {continue}
        fmt.Fprintf(&sb, "%d%c", i, b)
	}
	return sb.String()
}
