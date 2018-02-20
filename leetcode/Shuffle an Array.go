type Solution struct {
    raw []int
}


func Constructor(nums []int) Solution {
    return Solution{raw: nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.raw

}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    n := len(this.raw)
    words := make([]int, n)
    copy(words, this.raw)
    for i:=n; i>1; i-- {
        seed := rand.Intn(i)
        words[i-1], words[seed] = words[seed], words[i-1]
    }
    return words
}
