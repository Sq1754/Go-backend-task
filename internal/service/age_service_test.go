package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2000, 5, 10, 0, 0, 0, 0, time.UTC)
	age := CalculateAge(dob)

	if age <= 0 {
		t.Fatalf("expected positive age, got %d", age)
	}
}
