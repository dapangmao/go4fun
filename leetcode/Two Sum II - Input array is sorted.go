func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
    for i < j {
        var current = numbers[i] + numbers[j]
        if current == target {
            return []int{i+1, j+1}
        }
        if i < j && current < target {
            i++
        } else if i < j && current > target {
            j--
        }
    }
    return []int{}
}
