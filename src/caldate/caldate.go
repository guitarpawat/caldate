package caldate

import (
	"fmt"
	"strconv"
	"time"
	"math"
)

type Date struct {
	Date, Month, Year int
}

func UnitWeek(targetDate int) string {
	totalWeek := targetDate / 7
	totalDay := targetDate % 7
	if totalDay > 0 {
		resultHaveDay := strconv.Itoa(totalWeek) + " weeks and " + strconv.Itoa(totalDay) + " days"
		return resultHaveDay
	}
	resultNotHaveDay := strconv.Itoa(totalWeek) + " weeks"
	return resultNotHaveDay
}

func ResultDay(startDate, endDate Date) int {
	startTime := time.Date(startDate.Year, time.Month(startDate.Month), startDate.Date, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endDate.Year, time.Month(endDate.Month), endDate.Date, 0, 0, 0, 0, time.UTC)
	diff := endTime.Sub(startTime)
	//kkk := time
	return int(diff.Hours()/24) + 1
}

func ConvertToSecond(days int) uint64 {
	return uint64(days * 86400)
}

func ConvertToMin(second uint64) uint64 {
	return second / 60
}

type DetailStruct struct {
	Year, Month, Day int
}

func (a date) ResultDetailSameYear(endDate date) bool {
	return a.Year == endDate.Year
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
func FormatDateConverter(date Date) string {
	dateTime := time.Date(date.Year, time.Month(date.Month), date.Date, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %d %s %d", dateTime.Weekday().String(),
		dateTime.Day(), dateTime.Month().String(), dateTime.Year())
}

func CalPercent(days int) float64{
	percentile := math.Round(((float64(days) * 100) / 365)*100)/100
	return float64(percentile)
}
func toi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func NewDate(date, month, year string) Date {
	return Date{
		Date:  toi(date),
		Month: toi(month),
		Year:  toi(year),
	}
}
