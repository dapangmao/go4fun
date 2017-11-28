type MinStack struct {
    stack []int
    minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1]{
        this.minStack = append(this.minStack, x) 
    }
}

func (this *MinStack) Pop()  {
    var pop = this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    if pop == this.GetMin() {
       this.minStack = this.minStack[:len(this.minStack)-1] 
    }
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
