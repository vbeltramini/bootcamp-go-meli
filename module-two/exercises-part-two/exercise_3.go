package exercise

import "fmt"

func GetMonth() {
	fmt.Printf("\nExercise 3 - Months \n")

	var months = map[int]string{1: "Janeiro", 2: "Fevereiro", 3: "Março", 4: "Abril", 5: "Maio", 6: "Junho", 7: "Julho", 8: "Agosto", 9: "Setembro", 10: "Outubro", 11: "Novembro", 12: "Dezembro"}

	fmt.Println("Value passed: 3 | Month returned = ", months[3])
	fmt.Println("Value passed: 10 | Month returned = ", months[10])

}
