package exercise

import "fmt"

func FindAge() {
	fmt.Println("\nExercise 4 - Ages")

	var employess = map[string]int{"Benjamin": 20, "Manuel": 26, "Dario": 44, "Pedro": 30}

	getEmployessOlderThan21(employess)

	employess["Federico"] = 25

	delete(employess, "Pedro")

	fmt.Println("\nAge of Benjamin: ", employess["Benjamin"])

	getEveryEmployeeAge(employess)
}

func getEmployessOlderThan21(employess map[string]int) {
	fmt.Println("\nEmployers older than 21")

	for employeName, employeAge := range employess {
		if employeAge > 21 {
			fmt.Println(employeName, ": ", employeAge)
		}
	}

}

func getEveryEmployeeAge(employess map[string]int) {
	fmt.Println("\nEmployees ages: ")

	for employeName, employeAge := range employess {
		fmt.Println(employeName, ": ", employeAge)
	}
}
