package main

import (
	"fmt"

	exercises_part_one "github.com/vhbeltramini/bootcamp-go-meli/go-basics-part-two/exercises-part-one"
	exercises_part_two "github.com/vhbeltramini/bootcamp-go-meli/go-basics-part-two/exercises-part-two"
)

func main() {
	exercises_part_one.CalculateTaxe()
	exercises_part_one.CalcGrades(7, 7, 10, 10, 1, 8)
	exercises_part_one.CalculateSalary()
	exercises_part_one.SatisticCalcs()
	exercises_part_one.FeedAnimals()

	fmt.Println("\n--------------------------- Exercises Part Two -------------------------------------")

	exercises_part_two.PrintStudentDetails()

	fmt.Println()

	exercises_part_two.StoreExerciseMain()

}
