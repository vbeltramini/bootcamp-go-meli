package exercise

import (
	"errors"
	"fmt"
)

func CalculateSalary() {

	fmt.Println("Exercises 3 -> Calculate Salary")

	salary, _ := calculateBaseSalary("A", 190)
	fmt.Printf("salary 1 = %.2f\n", salary)
	salary, _ = calculateBaseSalary("B", 176)
	fmt.Printf("salary 2 = %.2f\n", salary)
	salary, _ = calculateBaseSalary("A", 172)
	fmt.Printf("salary 3 = %.2f\n", salary)
}

func calculateBaseSalary(category string, hours int) (float64, error) {

	switch category {
	case "A":
		return (float64(hours) * 3000) * getMultiplier(category, hours), nil
	case "B":
		return float64(hours) * 1500 * getMultiplier(category, hours), nil
	case "C":
		return (float64(hours) * 1000), nil
	default:
		return 0.0, errors.New("invalid category")
	}
}

func getMultiplier(category string, hours int) float64 {
	if hours > 160 {
		if category == "A" {
			return 1.5
		} else if category == "B" {
			return 1.2
		}
	}
	return 1
}
