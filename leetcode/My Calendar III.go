type MyCalendarThree struct {
    delta map[int]int
}


func Constructor() MyCalendarThree {
    return MyCalendarThree{delta: make(map[int]int)}
    
}


func (this *MyCalendarThree) Book(start int, end int) int {
    this.delta[start]++
    this.delta[end]--
    active, ans, i := 0, 0, 0
    keys := make([]int, len(this.delta))
    for k := range this.delta {
        keys[i] = k
        i++
    }
    sort.Ints(keys)
    for _, k := range keys {
        active += this.delta[k]
        if active > ans {ans = active}
    }
    return ans   
}
