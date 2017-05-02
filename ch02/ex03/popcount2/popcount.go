package popcount2

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var sum byte = 0
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(uint(i)*8))]
	}
	return int(sum)
}
