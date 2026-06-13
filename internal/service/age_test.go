package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Normal case",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: 36,
		},
		{
			name:     "Birthday not yet this year",
			dob:      time.Date(1990, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: 35,
		},
		{
			name:     "Birthday today",
			dob:      time.Date(time.Now().Year()-25, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
			expected: 25,
		},
		{
			name:     "Leap year DOB",
			dob:      time.Date(1996, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAge(tt.dob)
			if got != tt.expected {
				t.Errorf("CalculateAge(%v) = %d, want %d", tt.dob, got, tt.expected)
			}
		})
	}
}