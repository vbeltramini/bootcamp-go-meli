package exercises

import (
	"fmt"
	"math/rand"
)

func Ordering() {
	fmt.Println("\nExercise 4 Parte Two - Ordering")

	num := rand.Perm(100) // returns a slice of n ints

	fmt.Println("Array to order: ", num)
	fmt.Println("Ordered array:", short(num))
}

func short(itens []int) []int {

	for i := 1; i < len(itens); i++ {
		currentPosition := i
		fmt.Println(currentPosition)

		for currentPosition > 0 {
			if itens[currentPosition-1] > itens[currentPosition] {
				itens[currentPosition-1], itens[currentPosition] = itens[currentPosition], itens[currentPosition-1]
			}
			currentPosition -= 1
		}
	}

	return itens
}
