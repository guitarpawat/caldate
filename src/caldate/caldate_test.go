package caldate

import (
	"testing"
)

func Test_UnitWeek_Input_187_Should_Be_26_Weeks_And_5_Days(t *testing.T) {
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
