type MyHashSet struct {
    data map[int]bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{data: make(map[int]bool)}
}


func (this *MyHashSet) Add(key int)  {
    this.data[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    delete(this.data, key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return this.data[key]
}
