package exercises

import (
	"fmt"
	"sync"
)

type Service struct {
	name     string
	price    float64
	workTime int
}

type Maintenance struct {
	name  string
	price float64
}

var totalCost float64
var waitG sync.WaitGroup

func CalculatePriceMaintenance() {

	fmt.Println("\nExercise 3 Parte Two - Calculate Price (Part II)")

	var products = []Product{
		{Name: "Iphone 12", Price: 5000, Quantity: 10},
		{Name: "Macbook", Price: 10000, Quantity: 11},
		{Name: "USB C cable", Price: 19, Quantity: 12},
		{Name: "Mouse Pad", Price: 100, Quantity: 50},
	}

	var maintenance = []Maintenance{
		{name: "Fixing macbook", price: 1000},
		{name: "Installing an SSD", price: 300},
		{name: "Fixing iphone screen", price: 200},
		{name: "Cleaning Server room", price: 100},
	}

	var services = []Service{
		{name: "Service 1", price: 10, workTime: 60},
		{name: "Service 2", price: 1, workTime: 120},
		{name: "Service 3", price: 20, workTime: 10},
		{name: "Service 4", price: 15, workTime: 1000},
	}

	waitG.Add(3)
	go SumProducts(products)
	go SumMaintenance(maintenance)
	go SumServices(services)
	waitG.Wait()

	fmt.Println("The total cost of the services was ", totalCost)
}

func SumProducts(products []Product) {
	for _, product := range products {
		totalCost += product.Price * float64(product.Quantity)
	}
	waitG.Done()
}

func SumServices(services []Service) {
	for _, service := range services {
		if service.workTime < 30 {
			totalCost += service.price * 30
		}
		totalCost += service.price * float64(service.workTime)
	}
	waitG.Done()
}

func SumMaintenance(maintenance []Maintenance) {
	for _, maintenance := range maintenance {
		totalCost += maintenance.price
	}
	waitG.Done()
}
