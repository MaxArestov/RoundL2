package main

import (
	"testing"
	"time"
)

func TestTimeFormatting(t *testing.T) {
	testTime := time.Date(2023, time.April, 5, 15, 30, 45, 123000000, time.UTC)
	expected := "05.04.2023 15:30:45.123"
	result := testTime.Format("02.01.2006 15:04:05.000")
	if result != expected {
		t.Errorf("Неправильный формат времени, получено: %s, ожидалось: %s", result, expected)
	}
}
