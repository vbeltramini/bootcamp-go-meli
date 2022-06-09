package exercises

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string
	LastName string
	Email    string
	password string
	Age      int
}

type User interface {
	ChangeName(name, lastName string)
	ChangeAge(idade int)
	ChangeEmail(email string)
	ChangePassword(password string)
}

func UserManipulation() {
	fmt.Println("Exercise 1 Parte Two")

	var user user = user{Name: "Victor", LastName: "Beltramini", Email: "vhbeltramini@gmail.com", Age: 19, password: "098"}

	userPrint, _ := json.Marshal(user)
	fmt.Println("Initial User ", string(userPrint))

	user.ChangeName("Elon", "Musk")
	user.ChangeAge(30)
	user.ChangeEmail("elinho@gmail.com")
	user.ChangePassword("123456")

	userPrint, _ = json.Marshal(user)
	fmt.Println("Final User: ", string(userPrint))
}

func (u *user) ChangeName(name, lastName string) {
	log.Printf("Changing name from %s %s to %s %s", u.Name, u.LastName, name, lastName)
	(*u).Name = name
	(*u).LastName = lastName
}

func (u *user) ChangeAge(age int) {
	log.Printf("Changing age from %d to %d", u.Age, age)
	(*u).Age = age
}

func (u *user) ChangeEmail(email string) {
	log.Printf("Changing email from %s to %s", u.Email, email)
	(*u).Email = email
}

func (u *user) ChangePassword(password string) {
	log.Printf("Changing password from %s to %s", u.password, password)
	(*u).password = password
}
