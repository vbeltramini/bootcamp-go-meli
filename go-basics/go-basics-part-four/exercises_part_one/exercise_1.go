package exercises

import (
	"fmt"
)

var salary int

type customErrorMessage struct {
	message string
}

func (customError customErrorMessage) Error() string {
	return customError.message
}

var minimumSalaryError customErrorMessage
var mandatoryTaxesError customErrorMessage

func SalaryTaxesOne() {
	fmt.Println("\nExercise 1 Salary Tax #1")

	minimumSalaryError.message = "erro: O salário digitado não está dentro do valor mínimo"
	mandatoryTaxesError.message = "Necessário pagamento de imposto"

	salary = 1000

	if salary < 15000 {
		fmt.Println(minimumSalaryError.Error())
	} else {
		fmt.Println(mandatoryTaxesError.Error())
	}
}
