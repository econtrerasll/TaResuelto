package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Provider struct {
	ID   int32
	Name string
}

type Client struct {
	ID         int32
	Name       string
	ProviderID int
}

type Loan struct {
	ID         int32
	Balance    float64
	ProviderID int32
	ClientID   int32
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newClient(id int32, name string, idnumber int) {
	cliente := Client{ID: id, Name: name, ProviderID: idnumber}
	jclient, _ := json.Marshal(cliente)
	f, err := os.OpenFile("data/clients.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jclient))
}

func newLoan(id int32, balance float64, provid int32, cliID int32) {
	loan := Loan{ID: id, Balance: balance, ProviderID: provid, ClientID: cliID}
	jloan, _ := json.Marshal(loan)
	f, err := os.OpenFile("data/loans.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jloan))
}

func newProvider(id int32, name string) Provider {
	prov := Provider{ID: id, Name: name}
	jprov, _ := json.Marshal(prov)
	f, err := os.OpenFile("data/providers.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jprov))
	return prov
}

func listProviders() {
	r, err := ioutil.ReadFile("data/providers.json")
	check(err)
	fmt.Println(string(r))
}

func main() {

	newProvider(1, "Mateo")
	listProviders()
	newClient(1, "Miguel", 112)
	newLoan(1, 5000, 1, 112)
}
