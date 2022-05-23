package exercise

import (
	"errors"
	"fmt"
)

const (
	tarantulas = "tarantulas"
	hamsters   = "hamsters"
	dog        = "dog"
	cat        = "cat"
)

func tarantulasFoodFunc(quantity int) int {
	return quantity * 150
}

func hamstersFoodFunc(quantity int) int {
	return quantity * 250
}

func dogFoodFunc(quantity int) int {
	return quantity * 10000
}

func catFoodFunc(quantity int) int {
	return quantity * 5000
}

func animal(animalType string) (func(quantity int) int, error) {
	switch animalType {
	case tarantulas:
		return tarantulasFoodFunc, nil
	case hamsters:
		return hamstersFoodFunc, nil
	case cat:
		return catFoodFunc, nil
	case dog:
		return dogFoodFunc, nil
	default:
		return nil, errors.New("animal type not supported")
	}
}

func FeedAnimals() {
	fmt.Println("\nExercise 5 -> Feeding Animals...")
	fmt.Println("")

	hamstersFunc, _ := animal(hamsters)

	fmt.Printf("The food quantity for the dog is: %d \n", hamstersFunc(5))
	dogFunction, _ := animal(dog)

	fmt.Printf("The food quantity for the dog is: %d\n", dogFunction(5))
	catFunction, _ := animal(cat)

	fmt.Printf("The food quantity for the dog is: %d\n", catFunction(5))

}
