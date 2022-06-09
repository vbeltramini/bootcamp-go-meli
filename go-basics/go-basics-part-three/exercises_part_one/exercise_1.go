package exercise

import (
	"errors"
	"fmt"
	"os"
)

type product struct {
	id       int
	name     string
	quantity int
	price    float64
}

func gerarCsv(caminho string, products []product) error {
	if len(products) == 0 {
		return errors.New("you must provide at least one product")
	}
	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("error while opening file: %w", err)
	}
	defer file.Close()
	if _, err = file.WriteString(fmt.Sprintf("%10s,%10s,%10s,%10s", "id", "name", "quantity", "price\n")); err != nil {
		return fmt.Errorf("error while writing header: %w", err)
	}
	for _, p := range products {

		if _, err = file.WriteString(p.generateNewLineCSV()); err != nil {
			return fmt.Errorf("erro while writing a new line: %w", err)
		}
	}
	return nil
}

func (p product) generateNewLineCSV() string {
	return fmt.Sprintf("%10.d,%10s,%10.d,%10.2f\n", p.id, p.name, p.quantity, p.price)
}

func SaveProductsFile() {

	products := []product{
		{
			id:       1,
			name:     "USB C",
			quantity: 10,
			price:    19.99,
		},
		{
			id:       2,
			name:     "Banana",
			quantity: 100,
			price:    0.99,
		},
		{
			id:       3,
			name:     "Beer",
			quantity: 22,
			price:    7.99,
		},
		{
			id:       4,
			name:     "IDK",
			quantity: 30,
			price:    1.99,
		},
	}
	gerarCsv("productos.csv", products)
}
