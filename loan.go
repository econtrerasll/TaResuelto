package main

import (
	"encoding/json"
	"os"
)

//Loan is the main object that is provided and received in the platform
type Loan struct {
	ID        int32
	Balance   float64
	OwnerID   int32
	CreatorID int32
}

func newLoan(id int32, balance float64, creatorID int32, ownerID int32) {
	loan := Loan{ID: id, Balance: balance, CreatorID: creatorID, OwnerID: ownerID}
	jloan, _ := json.MarshalIndent(loan, "", " ")
	f, err := os.OpenFile("data/loans.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jloan))
}

func loadLoan(loanID int32) Loan {
	//Function to load saved loans from data resources
	l := Loan

	return l
}
