package caldate

func convertToHour(unitMin int) int {
	unitHour := unitMin / 60
	return unitHour
}

func convertToSecond(days int) int {
	return days * 86400
}
