package caldate

import (
	"testing"
)

func Test_ConvertToSecond_Input_182_Should_Be_15724800(t *testing.T) {
	expected := 15724800
	days := 182
	actual := convertToSecond(days)

	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
