package exercise

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Rg            string `json:"rg"`
	AdmissionDate string `json:"admission_date"`
}

func PrintStudentDetails() {
	fmt.Println("\nExercise One Part Two")
	fmt.Println("")

	student := Student{"Victor Hugo", "Grabowski Beltramini", "99.999.999", "16/04/2022"}
	resultJson, _ := json.Marshal(student)

	fmt.Println("Student :", string(resultJson))
}
