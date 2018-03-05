type RandomizedSet struct {
	location []int
	dict map[int]int

}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {return false}
	this.dict[val] = len(this.location)
	this.location = append(this.location, val)
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	loc, ok := this.dict[val]
	if !ok {return false}
	if loc > 0 {
		first := this.location[0]
		this.location[loc] = first
		this.dict[first] = loc
	}
	this.location = this.location[1:]
	delete(this.dict, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	loc := rand.Intn(len(this.location))
	return this.location[loc]
}
