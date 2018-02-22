type item struct {
	word string
	freq int
}

func topKFrequent(words []string, k int) []string {
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	var items []item
	for k, v := range dic {
		items = append(items, item{k, v})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].freq == items[j].freq {
			return items[i].word < items[j].word
		}
		return items[i].freq > items[j].freq
	})
	var res []string
	for i:=0; i<k; i++ {
		res = append(res, items[i].word)
	}
	return res
}
