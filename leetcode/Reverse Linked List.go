/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// more information can be found at TGPL 7.5.1
// **** an interface containing a Nil pointer is not nil ***
// such as new(ListNode)
// otherwise using *ListNode
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var prev *ListNode
    for head != nil {
        current := head.Next
        head.Next = prev
        prev = head
        head = current
    }
    return prev
}
