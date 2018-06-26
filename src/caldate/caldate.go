package caldate

func convertToSecond(days int) uint64 {
	return uint64(days * 86400)
}

func convertToMin(second uint64) uint64 {
	return second / 60
}
