package main

import (
	"encoding/json"
	"io/ioutil"
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
	//Function to load saved loans from data resources. This function should search in the existing saved data
	var l []Loan
	f, err := os.Open("data/loans.json") //read the data source
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f) //load the data
	if err != nil {
		panic(err)
	}
	//parse the data into an struct
	err = json.Unmarshal([]byte(data), &l)
	if err != nil {
		panic(err)
	}

	// search for the wanted object in the list
	for items := range l {
		if items == loanID {
			return items
		}
	}

}
