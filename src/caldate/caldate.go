package caldate

import (
	"fmt"
	"math"
	"strconv"
	"time"
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

func ResultDetail(startDate, endDate Date) Date {
	yearDiff := endDate.Year - startDate.Year
	monthDiff := endDate.Month - startDate.Month
	dayDiff := endDate.Date - startDate.Date + 1
	if monthDiff < 0 {
		yearDiff--
		monthDiff += 12
	}
	if dayDiff < 0 {
		monthDiff--
		dayDiff += 30
	}
	return Date{Year: yearDiff, Month: monthDiff, Date: dayDiff}
}

func (a Date) ResultDetailSameYear(endDate Date) bool {
	return a.Year == endDate.Year
}

func ConvertToHour(minute uint64) uint64 {
	return minute / 60

}

func FormatDateConverter(date Date) string {
	dateTime := time.Date(date.Year, time.Month(date.Month), date.Date, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %d %s %d", dateTime.Weekday().String(),
		dateTime.Day(), dateTime.Month().String(), dateTime.Year())
}

func CalPercent(days int) float64 {
	percentile := math.Round(((float64(days)*100)/365)*100) / 100
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

