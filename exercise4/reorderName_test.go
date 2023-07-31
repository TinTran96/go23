package ex4

import (
	"fmt"
	"testing"
)

// Test function with table-driven tests and subtests
func TestReorderNameVN(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"John", "Smith", "VN"}, "Smith John"},
		{[]string{"Emily", "Rose", "Watson", "US"}, "Emily Watson Rose"},
		{[]string{"Tran", "Tin", "US"}, "Tran Tin"},
		{[]string{"Azai", "Nagamasa", "JP"}, "Country Code JP Invalid"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := ReorderName(tc.input)
			if result != tc.expected {

				t.Errorf("For input '%s', expected %s, but got %s", tc.input, tc.expected, result)

			}
		})
	}
}
