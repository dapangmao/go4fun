type MinHeapInt []int

func (h MinHeapInt) Len() int {
	return len(h)
}

func (h MinHeapInt) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapInt) Peek() interface{} {
	return (*h)[0]
}

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
	length := len(*h)
	res := (*h)[length - 1]
	*h = (*h)[0 : length - 1]
	return res
}

type MaxHeapInt struct {
	MinHeapInt
}

func (h MaxHeapInt) Less(i, j int) bool {
	return h.MinHeapInt[i] > h.MinHeapInt[j]
}

type MedianFinder struct {
	maxHeap *MaxHeapInt
	minHeap *MinHeapInt
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minHeap := &MinHeapInt{}
	maxHeap := &MaxHeapInt{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{maxHeap, minHeap}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num <= this.maxHeap.Peek().(int) {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	} else if this.maxHeap.Len() - this.minHeap.Len() > 1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return (float64(this.maxHeap.Peek().(int)) + float64(this.minHeap.Peek().(int))) / 2.0
	} else {
		return float64(this.maxHeap.Peek().(int))
	}
}
