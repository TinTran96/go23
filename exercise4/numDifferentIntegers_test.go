package ex4

import (
	"fmt"
	"strconv"
	"testing"
)

func isEqual(a []int, b []int) bool {
	var aSize int = len(a)
	var bSize int = len(b)
	if aSize != bSize {
		return false
	}
	for i := 0; i < aSize; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func toStringArr(input []int) string {
	var resultStr string = ""
	var resultSize int = len(input)
	resultStr += "("
	for i := 0; i < resultSize; i++ {
		resultStr += strconv.Itoa(input[i])
		if i != (resultSize - 1) {
			resultStr += ", "
		}
	}
	resultStr += ")"
	return resultStr
}

// Test function with table-driven tests and subtests
func TestNumDifferentIntegers(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{"a123bc34d8ef34", []int{123, 34, 8}},
		{"a1b2c2d2fffff2", []int{1, 2}},
		{"avcdddddddddddddd", []int{}},
		{"", []int{}},
		{"a01bc0001c1dd00001", []int{1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := NumDifferentIntegers(tc.input)
			if !isEqual(result, tc.expected) {
				var resultStr string = toStringArr(result)
				var expectedStr string = toStringArr(tc.expected)

				t.Errorf("For input '%s', expected (%d)%s, but got (%d)%s", tc.input, len(tc.expected), expectedStr, len(result), resultStr)

			}
		})
	}
}
