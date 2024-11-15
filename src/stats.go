package main

import (
	"math"
)

func mean(array []int) float64 {
	sum := 0
	for _, num := range array {
		sum += num
	}
	result := float64(sum) / float64(len(array))

	return result
}

func meadian(numbers []int) float64 {
	n := len(numbers)
	if n%2 == 1 {
		return float64(numbers[n/2])
	}
	return (float64(numbers[n/2-1]) + float64(numbers[n/2])) / 2
}

func mode(numbers []int) int {
	maxCount := 0
	currentCount := 1
	mode := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
				mode = numbers[i-1]
			}
			currentCount = 1
		}
	}

	if currentCount > maxCount {
		mode = numbers[len(numbers)-1]
	}

	return mode
}

func standardDeviation(numbers []int, mean float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += math.Pow(float64(num)-mean, 2)
	}
	variance := sum / float64(len(numbers))
	return math.Sqrt(variance)
}
