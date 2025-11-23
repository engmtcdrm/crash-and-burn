package app

import "testing"

func TestSemVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"v1.2.3", "1.2.3"},
		{"1.2.3", "1.2.3"},
		{"v0.0.1", "0.0.1"},
		{"0.0.1", "0.0.1"},
		{"v1.2", "v1.2"},
		{"1.2", "1.2"},
		{"version1.2.3", "version1.2.3"},
		{"", ""},
	}

	for _, test := range tests {
		Version = test.input
		result := SemVersion()
		if result != test.expected {
			t.Errorf("SemVersion(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
