type MyCalendar struct {
    starts []int
    ends []int
    n int
    sdic map[int]int
}

func Constructor() MyCalendar {
    return MyCalendar{[]int{}, []int{}, 0, map[int]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
    if _, ok := this.sdic[start]; ok {return false}
    var i = sort.SearchInts(this.starts, start)
    if  (i < this.n && end > this.starts[i]) || (i > 0  && start < this.ends[i-1]) {return false}
    this.starts = append(this.starts[:i], append([]int{start}, this.starts[i:]...)...)
    this.ends = append(this.ends[:i], append([]int{end}, this.ends[i:]...)...)
    this.n++
    this.sdic[start] = 0
    return true
}
