type MyCalendar struct {
    starts []int
    ends []int
    n int
    sdic map[int]int
}

func Constructor() MyCalendar {
    return MyCalendar{[]int{}, []int{}, 0, map[int]int{}}
}


func (this *MyCalendar) add(arr *[]int, i, num int) []int {
    var res = make([]int, this.n+1)
    copy(res, (*arr)[:i])
    copy(res[i+1:], (*arr)[i:])
    res[i] = num
    return res
}

func (this *MyCalendar) Book(start int, end int) bool {
    if _, ok := this.sdic[start]; ok {return false}
    var i = sort.SearchInts(this.starts, start)
    if  (i < this.n && end > this.starts[i]) || (i > 0  && start < this.ends[i-1]) {return false}
    this.starts = this.add(&this.starts, i, start)
    this.ends = this.add(&this.ends, i, end)
    this.n++
    this.sdic[start] = 0
    return true
}
