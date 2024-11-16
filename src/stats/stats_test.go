package stats

import (
	"math"
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
			result := Mean(tt.numbers)
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
			result := Meadian(tt.numbers)
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
			result := Mode(tt.numbers)
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
			meanValue := Mean(tt.numbers)
			result := StandardDeviation(tt.numbers, meanValue)
			if math.Abs(result-tt.expected) > 1e-6 {
				t.Errorf("Expected %f, but got %f", tt.expected, result)
			}
		})
	}
}
