package api

import (
	"testing"
)

func Test__ToComma_Input_1810_Should_Be_1_Comma_810(t *testing.T) {
	expected := "1,810"
	targetDate := uint64(1810)
	actualDate := toComma(targetDate)
	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}

func Test__ToComma_Input_1234567_Should_Be_1_Comma_234_567(t *testing.T) {
	expected := "1,234,567"
	targetDate := uint64(1234567)
	actualDate := toComma(targetDate)
	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}

func Test__ToComma_Input_123_Should_Be_123(t *testing.T) {
	expected := "123"
	targetDate := uint64(123)
	actualDate := toComma(targetDate)
	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}