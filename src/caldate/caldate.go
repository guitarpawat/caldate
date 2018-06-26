package caldate

import (
	"strconv"
	"time"
)

func UnitWeek(targetDate int) string {
	totalWeek := targetDate / 7
	totalDay := targetDate % 7
	if totalDay > 0 {
		result1 := strconv.Itoa(totalWeek) + " weeks and " + strconv.Itoa(totalDay) + " days"
		return result1
	}
	result2 := strconv.Itoa(totalWeek) + " weeks"
	return result2
}

type date struct {
	Date, Month, Year int
}

const timeFormat = "2014-05-30"

func ResultDay(startDate, endDate date) int {
	startTime := time.Date(startDate.Year, time.Month(startDate.Month), startDate.Date, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endDate.Year, time.Month(endDate.Month), endDate.Date, 0, 0, 0, 0, time.UTC)
	diff := endTime.Sub(startTime)
	return int(diff.Hours()/24) + 1
}

func convertToSecond(days int) uint64 {
	return uint64(days * 86400)
}
