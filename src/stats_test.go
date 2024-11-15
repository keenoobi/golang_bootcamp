package main

import (
	"bufio"
	"errors"
	"math"
	"strings"
	"testing"
)

// проверяет правильность вычисления среднего значения
func TestMeanTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected float64
	}{
		{"Simple case", []int{1, 2, 3, 4, 5}, 3.0},
		{"With negative numbers", []int{-1, -2, -3, -4, -5}, -3.0},
		{"Single number", []int{10}, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mean(tt.numbers)
			if result != tt.expected {
				t.Errorf("Expected %f, but got %f", tt.expected, result)
			}
		})
	}
}

// проверяет правильность вычисления медианы
func TestMedianTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected float64
	}{
		{"Odd length", []int{1, 2, 3, 4, 5}, 3.0},
		{"Even length", []int{1, 2, 3, 4}, 2.5},
		{"Single number", []int{10}, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := meadian(tt.numbers)
			if result != tt.expected {
				t.Errorf("Expected %f, but got %f", tt.expected, result)
			}
		})
	}
}

// проверяет правильность вычисления моды
func TestModeTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Single mode", []int{1, 2, 2, 3, 4}, 2},
		{"Multiple modes", []int{1, 2, 2, 3, 3, 4}, 2}, // Возвращает первую встретившуюся моду
		{"Single number", []int{10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mode(tt.numbers)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, result)
			}
		})
	}
}

// проверяет правильность вычисления стандартного отклонения
func TestStandardDeviationTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected float64
	}{
		{"Simple case", []int{1, 2, 3, 4, 5}, 1.414214},
		{"With negative numbers", []int{-1, -2, -3, -4, -5}, 1.414214},
		{"Single number", []int{10}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meanValue := mean(tt.numbers)
			result := standardDeviation(tt.numbers, meanValue)
			if math.Abs(result-tt.expected) > 1e-6 {
				t.Errorf("Expected %f, but got %f", tt.expected, result)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []int
		expectedErr error
	}{
		{"Valid input", "1\n2\n3\n", []int{1, 2, 3}, nil},
		{"Invalid input", "1\nabc\n3\n", nil, errors.New("invalid input: the number is incorrect")},
		{"Empty input", "", []int{}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			numbers, err := parseInput(reader)
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
