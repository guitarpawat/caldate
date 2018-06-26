package caldate

import "testing"

func Test_ConvertToHour_Input_262080_Should_Be_4368(t *testing.T) {
	expected := 4368
	unitMin := 262080
	actual := convertToHour(unitMin)
	if actual != expected {
		t.Errorf("Expected %d but got %d ", expected, actual)
	}
}
