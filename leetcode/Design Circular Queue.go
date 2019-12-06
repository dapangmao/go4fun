type MyCircularQueue struct {
    data []int
    size int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{[]int{}, k}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if len(this.data) >= this.size {return false}
    this.data = append(this.data, value)
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if len(this.data) == 0 {return false}
    this.data = this.data[1:]
    return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if len(this.data) == 0 {return -1}
    return this.data[0]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if len(this.data) == 0 {return -1}
    return this.data[len(this.data)-1]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return len(this.data) == 0
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    return len(this.data) == this.size
}
