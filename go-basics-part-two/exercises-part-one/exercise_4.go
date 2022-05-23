package exercise

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func SatisticCalcs() {
	fmt.Println("\nExercise 4 ")

	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)
	min, _ := minFunc(2, 1, 3, 4, 10, 12)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 40, 30, 25)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 15, 3, 21, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}

func getFunc(funcName string) (func(values ...float64) (float64, error), error) {
	switch funcName {
	case maximum:
		return maxValue, nil
	case minimum:
		return minValue, nil
	case average:
		return averageValue, nil
	default:
		return nil, fmt.Errorf("invalid function type")
	}
}

func minValue(values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}

	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min, nil
}

func averageValue(values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}

	sum := 0.0

	for _, grade := range values {
		sum += grade
	}

	return sum / float64(len(values)), nil
}

func maxValue(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}

	var max float64
	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max, nil
}
