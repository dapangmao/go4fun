func merge(nums1 []int, m int, nums2 []int, n int)  {
    n--
    m--
    var j = n + m + 1
    for m >= 0 && n >= 0 {
        if nums1[m] >= nums2[n] {
            nums1[j] = nums1[m]
            m--
        } else {
            nums1[j] = nums2[n]
            n--
        }
        j--
    }
    for n >= 0 {
        nums1[j] = nums2[n]
        n--
        j--
    }
}
