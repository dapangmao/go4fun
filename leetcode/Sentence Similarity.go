// empty data structure is OK for a go map
func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
    dic := make(map[string][]string)
    for _, pair := range pairs {
        a, b := pair[0], pair[1]
        dic[a] = append(dic[a], b)
        dic[b] = append(dic[b], a)
    }
    n, m := len(words1), len(words2)
    if n != m {return false}
    for i:=0; i<n; i++ {
        w1, w2 := words1[i], words2[i]
        if w1 == w2 {goto next}
        for _, x := range dic[w1] {
            if x == w2 {goto next}
        }
        return false
        next:
    }
    return true
}
