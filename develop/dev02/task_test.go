package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{` 4`, `    `, false},
		{"a2b3c4d5", "aabbbccccddddd", false},
		{"ga43", "", true},
	}

	for _, tc := range testCases {
		result, err := unpackString(tc.input)
		if (err != nil) != tc.hasError {
			t.Errorf("unpackString(%s): ожидалась ошибка: %v, получена ошибка: %v", tc.input, tc.hasError, err)
		}
		if result != tc.expected {
			t.Errorf("unpackString(%s): ожидалось: %s, получено: %s", tc.input, tc.expected, result)
		}
	}
}
