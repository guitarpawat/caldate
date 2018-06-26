package caldate

import "strconv"

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
