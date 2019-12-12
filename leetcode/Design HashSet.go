type MyHashSet struct {
    data [1000001]bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    var res MyHashSet
    res.data = [1000001]bool{}
    return res
}


func (this *MyHashSet) Add(key int)  {
    this.data[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    this.data[key] = false
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return this.data[key]
}
