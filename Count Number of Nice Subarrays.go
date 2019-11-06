func numberOfSubarrays(nums []int, k int) int {
	magic := make([]int, 0)
	evenCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
		} else {
			magic = append(magic, evenCount+1)
			evenCount = 0
		}
	}
	magic = append(magic, evenCount+1)

	subArrays := 0
	for i := k; i < len(magic); i++ {
		subArrays += magic[i-k] * magic[i]
	}
	return subArrays
}
