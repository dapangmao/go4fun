type Skiplist struct {
    data [20001]int
}


func Constructor() Skiplist {
    return Skiplist{}
}


func (this *Skiplist) Search(target int) bool {
    return this.data[target] > 0
}


func (this *Skiplist) Add(num int)  {
    this.data[num]++
}


func (this *Skiplist) Erase(num int) bool {
    res := this.data[num]
    if res <= 0 {return false}
    this.data[num]--
    return true
}
