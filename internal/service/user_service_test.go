package service

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAge(t *testing.T) {
	now := time.Date(2025, 12, 16, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday has passed this year",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: 35,
		},
		{
			name:     "Birthday is yet to come this year",
			dob:      time.Date(1990, 12, 20, 0, 0, 0, 0, time.UTC),
			expected: 34,
		},
		{
			name:     "Birthday is today",
			dob:      time.Date(2000, 12, 16, 0, 0, 0, 0, time.UTC),
			expected: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob, now)
			assert.Equal(t, tt.expected, age)
		})
	}
}