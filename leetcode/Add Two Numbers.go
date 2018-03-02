/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func newNode(i int) *ListNode{
    var res = new(ListNode)
    res.Val = i
    return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var dummy = new(ListNode)
    var head = dummy
    var prev int
    for l1 != nil || l2 != nil {
        var current = prev
        if l1 != nil {
            current += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            current += l2.Val
            l2 = l2.Next
        }
        head.Next = newNode(current%10)
        prev = current/10
        head = head.Next
    }
    if prev > 0 {head.Next = newNode(prev)}
    return dummy.Next
}
