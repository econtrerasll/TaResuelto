package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type provider struct {
	ID   int32
	Name string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newProvider(id int32, name string) {
	prov := provider{ID: id, Name: name}
	jprov, _ := json.Marshal(prov)
	f, err := os.OpenFile("data/providers.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jprov))
}

func listProviders() {
	r, err := ioutil.ReadFile("data/providers.json")
	check(err)
	fmt.Println(r)

}

func main() {
	newProvider(1, "Mateo")
	newProvider(2, "Marcos")
	listProviders()
}
