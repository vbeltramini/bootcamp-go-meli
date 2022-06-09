package exercises

import (
	"errors"
	"fmt"
)

type customError error

var minSalaryError customError
var mandatoryTaxesErr customError

func SalaryTaxesTwo() {
	fmt.Println("\nExercise 2 Salary Tax #2")

	minSalaryError = errors.New("erro: o salário digitado não está dentro do valor mínimo")
	mandatoryTaxesErr = errors.New("necessário pagamento de imposto")

	salary = 15000

	if salary < 15000 {
		fmt.Println(minSalaryError.Error())
	} else {
		fmt.Println(mandatoryTaxesErr.Error())
	}
}
