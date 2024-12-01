package test

import (
	"bufio"
	"errors"
	pr "project/internal/parser"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []int
		expectedErr error
	}{
		{"Valid input", "1\n2\n3\n", []int{1, 2, 3}, nil},
		{"Invalid input", "1\nabc\n3\n", nil, errors.New("invalid input: the number is incorrect")},
		{"Invalid input max+1", "100001\n", nil, errors.New("invalid input: the number is incorrect")},
		{"Invalid input min-1", "-100001\n", nil, errors.New("invalid input: the number is incorrect")},
		{"Empty input", "", []int{}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			numbers, err := pr.ParseInput(reader)
			if err != nil && tt.expectedErr == nil {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && tt.expectedErr != nil {
				t.Errorf("Expected error: %v, but got nil", tt.expectedErr)
			} else if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedErr, err)
			}

			if !equal(numbers, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, numbers)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
