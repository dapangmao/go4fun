func numSubarrayBoundedMax(A []int, L int, R int) int {
    var res int
    
    factorial := func(n int) int {
        var res = 0
        for n > 0 {
            res += n
            n -= 1
        }
        return res
    }  
    
    compute := func(a []int) int {
        var res = factorial(len(a))
        var j = -1
        for i:=0; i<=len(a); i++{
            if i == len(a) || a[i] >= L {
                res -= factorial(i-j-1)
                j = i
            }
        }
        return res
    }
    
    j := -1
    for i:=0; i<=len(A); i++ {
        if i==len(A) || A[i] > R {
            if i>j+1 {
                res += compute(A[j+1:i])
            }
            j = i
        }
    }
    return res
}

