package popcount3

func PopCount(x uint64) (pc int) {
	for i := 0; i < 64; i++ {
		tmp := x >> uint(i)
		if tmp&1 == 1 {
			pc++
		}
	}
	return
}
