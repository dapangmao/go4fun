type MagicDictionary struct {
    dic map[int][]string
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
    return MagicDictionary{map[int][]string{}}
}


/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string)  {
    for _, word := range dict {
        var k = len(word)
        this.dic[k] = append(this.dic[k], word) 
    }
}

func (this MagicDictionary) isMagic(a, b string) bool {
    var count int
    for i := range a {
        if a[i] != b[i] {count++}
        if count > 1 {return false}
    }
    return count == 1
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
    var target = this.dic[len(word)]
    for _, x := range target {
        if this.isMagic(x, word) {return true}
    }
    return false
}
