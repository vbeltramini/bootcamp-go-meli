package exercises

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const minID int = 100000
const maxID int = 999999

type Client struct {
	Archive   int
	Name      string
	LastName  string
	RG        string
	Telephone string
	Address   string
}

func erro() {
	err := recover()

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func RegisteringClients() {
	defer erro()

	user := newClient("Victor", "Hugo", "999.999-00", "912345678", "Rua Tiradentes")
	fmt.Println(user)
}

func newClient(name, lastName, rg string, telephone string, address string) Client {
	if name == "" || rg == "" || telephone == "" || address == "" {
		panic("Please enter all values")
	}

	newClient := Client{23212, name, lastName, rg, telephone, address}

	clients := abrirCsv("/Users/vbeltramini/Documents/bootcamp/go-basics/bootcamp-go-meli/go-basics-part-four/costumers.csv")
	for _, client := range clients {
		temp := client.Archive
		client.Archive = 0

		if client == newClient {
			panic("o client já existe")
		}
		client.Archive = temp
	}

	newClient.Archive = generateNewId()

	return newClient
}

func generateNewId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxID-minID) + minID
}

func abrirCsv(local string) []Client {
	arquivo, err := os.Open(local)
	if err != nil {
		panic("caminho incorreto ou arquivo inexistente")
	}
	defer arquivo.Close()

	linesCsv, err := csv.NewReader(arquivo).ReadAll()
	if err != nil {
		panic("separador incorreto ou arquivo csv inválido")
	}

	var listClients []Client
	for _, client := range linesCsv {
		id, _ := strconv.Atoi(client[0])
		user := Client{id, client[1], client[2], client[3], client[4], client[5]}

		listClients = append(listClients, user)
	}

	return listClients
}
