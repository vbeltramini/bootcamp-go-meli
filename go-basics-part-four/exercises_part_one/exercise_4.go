package exercises

import (
	"errors"
	"fmt"
)

func CalculateEmployeeSalary() {
	fmt.Println("\nExercise 4 Salary Tax #4")

	valor, err := CalculateSalary(85, 1000)

	if err != nil {
		err = fmt.Errorf("%s | informed worked hours = %d", err, 85)
		fmt.Println(err)
		return
	}

	fmt.Println("Salary: ", valor)
}

func CalculateSalary(workedHours int, hourPrice float64) (salario float64, err error) {
	salary := float64(workedHours) * hourPrice

	if workedHours > 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	if salary >= 20000 {
		salary -= salary * 0.10
	}
	return salary, nil
}
