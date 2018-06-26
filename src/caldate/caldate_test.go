package caldate

import "testing"

func Test_ResultDay_Input_StartDate_4_1_2018_EndDate_4_7_2018_Should_Be_182(t *testing.T) {
	startDate := date{Date: 4, Month: 1, Year: 2018}
	endDate := date{Date: 4, Month: 7, Year: 2018}
	expected := 182
	result := ResultDay(startDate, endDate)
	if expected != result {
		t.Errorf("expected %d but get %d", expected, result)
	}
}
