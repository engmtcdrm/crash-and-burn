package settings

import (
	"testing"
)

func TestFailureRCs_Set(t *testing.T) {
	tests := []struct {
		input       string
		expectedErr bool
	}{
		{"1,50", false},
		{"256,50", true},
		{"1,101", true},
		{"a,1", true},
		{"1,a", true},
		{"1.1,1", true},
		{"1,1.1", true},
		{"true,1", true},
		{"1,true", true},
	}

	for _, test := range tests {
		var e FailureRCs
		err := e.Set(test.input)
		if (err != nil) != test.expectedErr {
			t.Errorf("Set(%s) error = %v, expectedErr %v", test.input, err, test.expectedErr)
		}
	}
}

func TestFailureRCs_TotalPct(t *testing.T) {
	e := FailureRCs{
		{RC: 1, Pct: 30},
		{RC: 2, Pct: 40},
		{RC: 3, Pct: 20},
	}

	expectedTotal := 90
	if total := e.TotalPct(); total != expectedTotal {
		t.Errorf("TotalPct() = %d, expected %d", total, expectedTotal)
	}
}

func TestFailureRCs_exists(t *testing.T) {
	e := FailureRCs{
		{RC: 1, Pct: 30},
		{RC: 2, Pct: 40},
	}

	tests := []struct {
		rc       int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, false},
	}

	for _, test := range tests {
		if exists := e.exists(test.rc); exists != test.expected {
			t.Errorf("exists(%d) = %v, expected %v", test.rc, exists, test.expected)
		}
	}
}
