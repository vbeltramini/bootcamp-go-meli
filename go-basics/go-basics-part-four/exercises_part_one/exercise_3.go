package exercises

import (
	"errors"
	"fmt"
)

type customErr error

var minSalaryErr customErr
var mandatoryTaxErr customErr

func SalaryTaxesThree() {
	fmt.Println("\nExercise 3 Salary Tax #3")

	minSalaryErr = errors.New("erro: o mínimo tributável é 15000")
	mandatoryTaxErr = errors.New("necessário pagamento de imposto")

	salary = 1000
	if salary < 15000 {
		err := fmt.Errorf("%s e o salário informado é: R$%d", minSalaryErr.Error(), salary)
		fmt.Println(err)
	} else {
		fmt.Println(mandatoryTaxErr.Error())
	}
}
