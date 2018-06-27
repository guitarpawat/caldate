package caldate

import (
	"testing"
)

func Test_UnitWeek_Input_187_Should_Be_26_Weeks(t *testing.T) {
	expected := "26 weeks"
	targetDate := 182
	actualDate := UnitWeek(targetDate)
	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}
func Test_UnitWeek_Input_182_Should_Be_26_Weeks_And_5_Days(t *testing.T) {
	expected := "26 weeks and 5 days"
	targetDate := 187
	actualDate := UnitWeek(targetDate)
	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}
func Test_ConvertToSecond_Input_182_Should_Be_15724800(t *testing.T) {
	expected := uint64(15724800)
	days := 182
	actual := ConvertToSecond(days)

	if actual != expected {
		t.Errorf("expected %d but got %d", uint64(expected), uint64(actual))
	}
}

func Test_ConvertToMin_Input_15724806_Should_Be_262080(t *testing.T) {
	expected := uint64(262080)
	second := uint64(15724806)
	actual := uint64(ConvertToMin(second))

	if actual != expected {
		t.Errorf("expected %d but got %d", uint64(expected), uint64(actual))
	}
}

func Test_ResultDay_Input_StartDate_4_1_2018_EndDate_4_7_2018_Should_Be_182(t *testing.T) {
	startDate := Date{Date: 4, Month: 1, Year: 2018}
	endDate := Date{Date: 4, Month: 7, Year: 2018}
	expected := 182
	result := ResultDay(startDate, endDate)
	if expected != result {
		t.Errorf("expected %d but get %d", expected, result)
	}
}

func Test_FormatDateConverter_Input_4_1_2018_Should_Be_Thursday_4_January_2018(t *testing.T) {
	date := Date{Date: 4, Month: 1, Year: 2018}
	expected := "Thursday, 4 January 2018"
	result := FormatDateConverter(date)
	if expected != result {
		t.Errorf("expected %s but get %s", expected, result)
	}
}
func Test_CalPercent_Input_182_Should_Be_49_dot_86(t *testing.T) {
	days := 182
	expected := 49.86
	result := CalPercent(days)
	if expected != result {
		t.Errorf("expected %.2f but get %.2f", expected, result)
	}
}
func Test_CalPercent_Input_7248_Should_Be_1985_dot_75(t *testing.T) {
	days := 7248
	expected := 1985.75
	result := CalPercent(days)
	if expected != result {
		t.Errorf("expected %.2f but get %.2f", expected, result)
	}
}

func Test_ConvertToHour_Input_262080_Should_Be_4368(t *testing.T) {
	var minute uint64 = 262080
	var expected uint64 = 4368
	result := ConvertToHour(minute)
	if expected != result {
		t.Errorf("expected %d but get %d", expected, result)
	}
}
