func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) {return false}
    visited, dic := make(map[string]int), make(map[string][]string)
    count := 1
    for _, pair := range pairs {
        a, b := pair[0], pair[1]
        visited[a] = -count
        count++
        visited[b] = -count
        count++
        dic[a] = append(dic[a], b)
        dic[b] = append(dic[b], a)
    }
    var group int
    for k, v := range visited {
        if v > 0 {continue}
        group++
        queue := []string{k}
        for len(queue) > 0 {
            first := queue[0]
            queue = queue[1:]
            visited[first] = group
            for _, x := range dic[first] {
                if visited[x] < 0 {
                    queue = append(queue, x)
                    visited[x] = group
                }
            }
        }
    }
    for i:=0; i<len(words1); i++ {
        w1, w2 := words1[i], words2[i]
        if visited[w1] != visited[w2] {return false}
    }
    return true
}
