package main

import (
	"testing"
)

func TestGrepFile(t *testing.T) {
	tests := []struct {
		name           string
		pattern        string
		config         *grepConfig
		expectedCount  int
		expectedOutput []string
	}{
		{
			name:    "Basic match",
			pattern: "test",
			config: &grepConfig{
				0,
				0,
				0,
				false,
				false,
				false,
				false,
				false,
			},
			expectedCount: 4,
			expectedOutput: []string{
				"This is a test file.",
				"This is another test line.",
				"Lines with 'test' should be matched by grep.",
				"This is a test file.",
			},
		},
		{
			name:    "After match",
			pattern: "should",
			config: &grepConfig{
				2,
				0,
				0,
				false,
				false,
				false,
				false,
				false,
			},
			expectedCount: 1,
			expectedOutput: []string{
				"Lines with 'test' should be matched by grep.",
				"This is the last line.",
				"This is a test file.",
			},
		},
		{
			name:    "Context match ignore case",
			pattern: "should",
			config: &grepConfig{
				0,
				0,
				2,
				false,
				true,
				false,
				false,
				false,
			},
			expectedCount: 1,
			expectedOutput: []string{
				"this is another test line.",
				"grep is a powerful utility.",
				"lines with 'test' should be matched by grep.",
				"This is the last line.",
				"This is a test file.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, count, err := GrepFile("file.txt", tt.pattern, tt.config)
			if err != nil {
				t.Errorf("GrepFile returned an error: %v", err)
			}
			if count != tt.expectedCount {
				t.Errorf("Expected %d matches, got %d", tt.expectedCount, count)
			}
			if len(lines) != len(tt.expectedOutput) {
				t.Fatalf("Expected %d lines, got %d", len(tt.expectedOutput), len(lines))
			}
			for i, line := range lines {
				if line != tt.expectedOutput[i] {
					t.Errorf("Expected line %d to be '%s', got '%s'", i, tt.expectedOutput[i], line)
				}
			}
		})
	}
}
