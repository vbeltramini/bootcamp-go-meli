package exercise

import (
	"errors"
	"fmt"
	"log"
)

func calc(grades []float64) (float64, error) {

	sum := 0.0

	for _, grade := range grades {

		if grade < 0 {
			return grade, errors.New("grade must be positive")
		}

		sum += grade
	}

	return sum / float64(len(grades)), nil
}

func CalcGrades(grades ...float64) {
	fmt.Println("\nExercise 2 -> Calculate average grade")

	firstAveraged, err := calc(grades)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Average : %.2f\n", firstAveraged)
}
