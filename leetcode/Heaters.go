func findRadius(houses []int, heaters []int) int {
  sort.Ints(heaters)
	var radius = 0
	for _, x := range houses {
		var j = sort.SearchInts(heaters, x)
		if j < len(heaters) && x == heaters[j] {continue}
		var a, b = 1<<32, 1<<32
		if j-1 >= 0 {a = x - heaters[j-1] }
		if j < len(heaters) {b = heaters[j] - x}
		var current = a
		if b < a {current = b}
		if current > radius {radius = current}
	}
	return radius
}
