package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		lines  []string
		config *flagConfig
		want   []string
	}{
		{
			name:   "TestNumericSort",
			lines:  []string{"10", "2", "1"},
			config: &flagConfig{numericSort: true},
			want:   []string{"1", "2", "10"},
		},
		{
			name:   "MonthSort",
			lines:  []string{"February", "January", "March"},
			config: &flagConfig{monthSort: true},
			want:   []string{"January", "February", "March"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.lines, tt.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSuffixNumber(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    int
		wantErr bool
	}{
		{
			name: "TestK",
			s:    "1K",
			want: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSuffixNumber(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSuffixNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseSuffixNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
