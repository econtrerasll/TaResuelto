package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Client describe the entity that received the loans
type Client struct {
	ID         int32
	Name       string
	ProviderID int32
}

//Loan is the main object that is provided and received in the platform
type Loan struct {
	ID        int32
	Balance   float64
	OwnerID   int32
	CreatorID int32
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newClient(id int32, name string, provID int32) {
	cli := Client{ID: id, Name: name, ProviderID: provID}
	jcli, _ := json.Marshal(cli)
	f, err := os.OpenFile("data/clients.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jcli))
}

func newLoan(id int32, balance float64, creatorID int32, ownerID int32) {
	loan := Loan{ID: id, Balance: balance, CreatorID: creatorID, OwnerID: ownerID}
	//jloan, _ := json.Marshal(loan)
	jloan, _ := json.MarshalIndent(loan, "", " ")
	f, err := os.OpenFile("data/loans.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jloan))
}

func listProviders() {
	r, err := ioutil.ReadFile("data/providers.json")
	check(err)
	fmt.Println(string(r))
}

func main() {
	giveLoan(8000, 1, 1)
}
