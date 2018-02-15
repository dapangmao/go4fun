type MapSum struct {
    dic map[string]int
    lastInsert map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{map[string]int{}, map[string]int{}}
}

func (this *MapSum) Insert(key string, val int)  {
    lastInsertVal, ok := this.lastInsert[key]
    for i:=1; i<=len(key); i++ {
        var current = key[:i]
        this.dic[current] += val
        if ok {this.dic[current] -= lastInsertVal}
    }
    this.lastInsert[key] = val
}


func (this *MapSum) Sum(prefix string) int {
    return this.dic[prefix]
}
