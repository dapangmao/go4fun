/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := new(ListNode)
    dummy.Next = head
    prev := dummy
    for head != nil {
        if head.Val == val {
            prev.Next = head.Next
        } else {
            prev = head
        }
        head = head.Next
    }
    return dummy.Next
}
