package utils

func sum(i uint64) (sum int) {
	b64 := uint64(10)
	for ; i > 0; i /= b64 {
		sum += int(i % b64)
	}
	return
}
func total(i int) (total int) {
	if i > 9 && i != 11 {
		total = sum(uint64(i))
	} else {
		total = i
	}
	return
}

// CalculateLifePath sums each individual date portion, then
// sums again if a date portion is greater than 9 and does not equal 11.
func CalculateLifePath(d uint64, m uint64, y uint64) int {
	day := total(sum(d))
	month := total(sum(m))
	year := total(sum(y))
	return day + month + year
}
