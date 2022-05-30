package exercises

import (
	"fmt"
	"os"
)

func readTXT(filepath string) error {
	_, err := os.Open(filepath)
	defer recover()
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	return nil
}

func ClientData() {
	fmt.Println("\nExercise 1 | Part two : Client Data ")

	readTXT("customers.txt")
}
