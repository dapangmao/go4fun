type LRUCache struct {
    capacity int
    cache map[int]int
    order []int
    pos map[int]int
}

func Constructor(capacity int) LRUCache {
    var c = make(map[int]int)
    var o = []int{}
    var p = make(map[int]int)
    return LRUCache{capacity: capacity, cache: c, order: o, pos: p}
}

func (this *LRUCache) Get(key int) int {
    if val, ok := this.cache[key]; ok {
        var p = this.pos[key]
        this.order = append(this.order[:p], this.order[p+1:]...)
        this.order = append(this.order, key)
        this.pos[key] = len(this.cache) - 1
        return val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    var current = this.Get(key)
    if current != -1 {
        this.cache[key] = value
        return
    }
    if len(this.cache) == this.capacity {
        var current = this.order[0]
        delete(this.cache, current)
        this.order = this.order[1:]
    }
    this.order = append(this.order, key)
    this.cache[key] = value
    this.pos[key] = len(this.cache) - 1
}
