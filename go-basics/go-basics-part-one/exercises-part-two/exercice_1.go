package exercise

import (
	"fmt"
)

func LetterOfAWord() {
	fmt.Println("\nExercise number 1 - Part Two")

	word := "Teste"

	wordlength := len([]rune(word))

	fmt.Println("Total of Words : ", wordlength, "\n ")
	for i := 0; i < wordlength; i++ {
		fmt.Println(string([]rune(word)[i]))
	}

}
