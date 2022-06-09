package exercise

import "fmt"

const (
	small  = "small"
	medium = "medium"
	large  = "large"
)

type product struct {
	Name  string
	Price float64
	Size  string
}

type Product interface {
	CalculatePrice() float64
}

type store struct {
	products []Product
}

type Ecommerce interface {
	Total() float64
	Add(p product)
}

func NewProduct(name string, price float64, size string) Product {
	newProduct := product{name, price, size}
	return &newProduct
}

func (p *product) CalculatePrice() float64 {
	switch p.Size {
	case medium:
		return (p.Price * 0.03) + p.Price
	case large:
		return (p.Price * 0.06) + p.Price + 2500
	default:
		return p.Price
	}
}

func (s *store) Total() float64 {
	total := 0.0
	for _, product := range s.products {
		total += product.CalculatePrice()
	}
	return total
}

func (s *store) Add(p Product) {
	s.products = append(s.products, p)
}

func StoreExerciseMain() {
	selectedProducts := []Product{
		NewProduct("PC", 100, large),
		NewProduct("PC", 100, medium),
		NewProduct("PC", 100, large),
	}

	myStore := store{}

	for _, product := range selectedProducts {
		myStore.Add(product)
	}

	fmt.Println(myStore.Total())
}
