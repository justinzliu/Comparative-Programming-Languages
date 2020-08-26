package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions
func Hailstone(x uint) uint {
	//even
	if (x%2 == 0) {
		return x/2
	}
	return ((3*x)+1)
}

func HailstoneSequenceAppend(x uint) []uint {
	start := []uint{}
	curr := x
	for curr > 1 {
		start = append(start,curr)
		curr = Hailstone(curr)
	}
	start = append(start,1)
	return start
}

func HailstoneLen(x uint) uint {
	curr := x
	var len uint = 1
	for curr > 1 {
		curr = Hailstone(curr)
		len++
	}
	return len
}

func HailstoneSequenceAllocate(x uint) []uint {
	len := HailstoneLen(x)
	start := make([]uint, len, len)
	curr := x
	it := 0
	for curr > 1 {
		start[it] = curr
		curr = Hailstone(curr)
		it++
	}
	start[len-1] = 1
	return start
}