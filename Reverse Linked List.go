/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    prev := new(ListNode).Next
    for head != nil {
        current := head.Next
        head.Next = prev
        prev = head
        head = current
    }
    return prev
}
