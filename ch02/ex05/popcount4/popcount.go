package popcount4

func PopCount(x uint64) (pc int) {
	for x > 0 {
		cleared := x & (x - 1)
		if x != cleared {
			pc++
		}
		x = cleared
	}
	return
}
