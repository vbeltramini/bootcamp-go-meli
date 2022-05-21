package exercise

import "fmt"

func CalculateTaxe() {
	fmt.Println("\nExericise One - Salary Taxes")

	fmt.Println("Case one: Salary = 48000")
	fmt.Println("Taxes = ", CalculateSalaryTaxes(48000))

	fmt.Println("Case one: Salary = 55000")
	fmt.Println("Taxes = ", CalculateSalaryTaxes(55000))

	fmt.Println("Case one: Salary = 155000")

	fmt.Println("Taxes = ", CalculateSalaryTaxes(155000))

}

func CalculateSalaryTaxes(salary float32) float32 {
	switch true {
	case salary >= 150000:
		return salary * 0.027
	case salary >= 50000:
		return salary * 0.017
	default:
		return 0
	}
}
