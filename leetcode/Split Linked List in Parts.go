func splitListToParts(root *ListNode, k int) []*ListNode {
    res := make([]*ListNode, k)
    var n int
    head := root
    for head != nil {
        n++
        head = head.Next
    }
    size, plus := n/k, n%k
    head = root
    var prev = new(ListNode)
    var i, j int
    for head != nil {
        if i == 0 || (plus == 0 && i == size) || (plus > 0 && i == size+1) {
            res[j] = head
            j++
            prev.Next = nil
            if i != 0 && plus > 0 {plus--}
            i = 0
        }
        prev = head
        head = head.Next
        i++
    }
    return res
}
