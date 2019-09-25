func isAlienSorted(words []string, order string) bool {
   
    dict := make(map[byte]int)
    for i, c := range []byte(order) {
        dict[c] = i
    }
    
    isTwoWordsSorted := func(a, b string) bool {
        i := 0
        for {
            if i == len(a) {return true}
            if i == len(b) {return false}
            if dict[a[i]] < dict[b[i]] {return true}
            if dict[a[i]] > dict[b[i]] {return false}
            i++
        }
        return true
    }
    
    for i:=1; i<len(words); i++ {
        if !isTwoWordsSorted(words[i-1], words[i]) {return false}
    }
    return true
}
