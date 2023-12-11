package main

import (
	"testing"
)

func TestCutLines(t *testing.T) {
	tests := []struct {
		name           string
		line           string
		config         cutConfig
		expectedResult string
	}{
		{
			name:           "Basic test with default delimiter",
			line:           "a\tb\tc",
			config:         cutConfig{fieldsFlag: "1", delimiterFlag: "\t"},
			expectedResult: "a",
		},
		{
			name:           "Test with custom delimiter",
			line:           "a,b,c",
			config:         cutConfig{fieldsFlag: "2", delimiterFlag: ","},
			expectedResult: "b",
		},
		{
			name:           "Test with separated flag",
			line:           "a\tb\tc",
			config:         cutConfig{fieldsFlag: "1", delimiterFlag: "\t", separatedFlag: true},
			expectedResult: "a",
		},
		{
			name:           "Test without delimiter in line when separated flag is set",
			line:           "abc",
			config:         cutConfig{fieldsFlag: "1", delimiterFlag: "\t", separatedFlag: true},
			expectedResult: "",
		},
		{
			name:           "Test with multiple fields",
			line:           "1,2,3,4,5",
			config:         cutConfig{fieldsFlag: "1,3,5", delimiterFlag: ","},
			expectedResult: "1,3,5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CutLines(&tt.config, tt.line)
			if result != tt.expectedResult {
				t.Errorf("TestCutLines(%s): expected %s, got %s", tt.name, tt.expectedResult, result)
			}
		})
	}
}
