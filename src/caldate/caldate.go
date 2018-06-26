package caldate

func convertToHour(unitMin int) int {
	unitHour := unitMin / 60
	return unitHour
}
func convertToSecond(days int) uint64 {
	return uint64(days * 86400)
}
