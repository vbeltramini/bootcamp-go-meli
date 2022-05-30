package exercises

import (
	"encoding/json"
	"fmt"
)

type EcommerceUser struct {
	Name     string
	Products []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func Ecommerce() {
	fmt.Println("\nExercise 2 Parte Two - Ecommerce")

	newUser := EcommerceUser{Name: "Victor"}

	addProduct(&newUser, newProduct("Macbook", 5000, 1))
	addProduct(&newUser, newProduct("Iphone 30", 15000, 1))
	addProduct(&newUser, newProduct("Mesa", 500, 1))

	fmt.Printf("\n Products : %v of user : %s", newUser.Products, newUser.Name)

	deleteProducts(&newUser)
}

func newProduct(name string, price float64, quantity int) Product {
	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func addProduct(user *EcommerceUser, product Product) {
	user.Products = append(user.Products, product)
	productInfo, _ := json.Marshal(product)
	fmt.Printf("Product : %s add to user : %s \n", string(productInfo), user.Name)
}

func deleteProducts(user *EcommerceUser) {
	user.Products = nil
	fmt.Printf("\nAll products from user %s have been deleted \n", user.Name)
}
