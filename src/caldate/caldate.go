package caldate

import (
	"time"
)

type Date struct {
	Date, Month, Year int
}

const timeFormat = "2014-05-30"

func ResultDay(startDate, endDate Date) int {
	startTime := time.Date(startDate.Year, time.Month(startDate.Month), startDate.Date, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endDate.Year, time.Month(endDate.Month), endDate.Date, 0, 0, 0, 0, time.UTC)
	diff := endTime.Sub(startTime)
	return int(diff.Hours()/24) + 1
}

func convertToSecond(days int) uint64 {
	return uint64(days * 86400)
}
