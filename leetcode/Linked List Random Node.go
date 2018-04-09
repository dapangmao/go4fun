/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    ptrs []*ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    var ptrs []*ListNode
    for head != nil {
        ptrs = append(ptrs, head)
        head = head.Next
    }
    return Solution{ptrs}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    // should be Intn instead of Int
    var current = rand.Intn(len(this.ptrs))
    return this.ptrs[current].Val
}
