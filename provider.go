package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

// Provider described the entity that give the loans to the clients
type Provider struct {
	ID   int32
	Name string
}

func newProvider(id int32, name string) {
	prov := Provider{ID: id, Name: name}
	jprov, _ := json.Marshal(prov)
	f, err := os.OpenFile("data/providers.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jprov))
}

func giveLoan(amount float64, providerID int32, clientID int32) {
	// Todo for later, add here logic to generate the id for next loan based on the last number id given in the list
	// of loans.
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100000
	newLoan(int32(rand.Intn(max-min)), amount, providerID, clientID)
}

func collectLoan(loanID int32, amount float64) {
	//find the loan by ID, then reduce it balance by amount collected
	loadLoan(loanID)
}
