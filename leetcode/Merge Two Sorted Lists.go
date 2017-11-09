/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode) 
    dummy := head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val{
            dummy.Next = l1
            l1 = l1.Next
        } else {
            dummy.Next = l2
            l2 = l2.Next
        }
        dummy = dummy.Next
    }
    if l1 != nil {
        dummy.Next = l1
    } else {
        dummy.Next = l2
    }
    return head.Next
}
