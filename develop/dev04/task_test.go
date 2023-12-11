package main

import (
	"reflect"
	"testing"
)

// Функция тестирования findAnagrams
func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "Basic Test",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"листок": []string{"листок", "слиток", "столик"},
				"пятак":  []string{"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:     "Empty List",
			words:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "Single Element Groups",
			words:    []string{"одно", "слово", "тест"},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findAnagrams(test.words)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test '%s' failed: expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}
