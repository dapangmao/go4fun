/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// more information can be found at TGPL 7.5.1
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    // (1) untyped nil does not work
    // prev := nil 
    // (2) This works
    // prev := new(ListNode).Next
    // (3) Best approach
    var prev *ListNode
    for head != nil {
        current := head.Next
        head.Next = prev
        prev = head
        head = current
    }
    return prev
}
