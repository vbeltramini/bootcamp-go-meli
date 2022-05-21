package exercise

import "fmt"

func BasicLoanValidator() {
	fmt.Println("\nExercise 2: Loan")

	var age, salary, yearsInActive uint32 = 19, 100001, 2

	println("\n Case one => age: ", age, ", salary: ", salary, ", active working years: ", yearsInActive)

	evaluate(age, salary, yearsInActive)

	age, salary, yearsInActive = 25, 98000, 2

	println("\n Case one => age: ", age, ", salary: ", salary, ", active working years: ", yearsInActive)

	evaluate(age, salary, yearsInActive)

	age, salary, yearsInActive = 25, 198000, 0

	println("\n Case one => age: ", age, ", salary: ", salary, ", active working years: ", yearsInActive)

	evaluate(age, salary, yearsInActive)

}

func evaluate(age, salary, yearsInActive uint32) {
	fmt.Println()
	switch false {
	case age > 22:
		fmt.Println("You must be older than 22, your age now is: ", age)
	case salary > 100000:
		fmt.Println("You salary must be greater than 100.000, you salary now is : ", salary)
	case yearsInActive > 1:
		fmt.Println("You must have to be an active worker for more than a year, now are just ", yearsInActive, " working here")
	}
}
