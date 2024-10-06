package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		slice    []int
		value    int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{}, 1, false},
		{[]int{1, 1, 1, 1}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
	}

	for _, test := range tests {
		result := contains(test.slice, test.value)
		if result != test.expected {
			t.Errorf("contains(%v, %d) = %v; want %v", test.slice, test.value, result, test.expected)
		}
	}
}

func TestPluralize(t *testing.T) {
	tests := []struct {
		count    int
		expected string
	}{
		{-2, "s"},
		{-1, ""},
		{0, "s"},
		{1, ""},
		{2, "s"},
	}

	for _, test := range tests {
		result := pluralize(test.count)
		if result != test.expected {
			t.Errorf("pluralize(%d) = %s; want %s", test.count, result, test.expected)
		}
	}
}
