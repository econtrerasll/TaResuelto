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
	providerID int
}

func newClient(id int32, name string, provider int) {
	cliente := Client{ID: id, Name: name, providerID: provider}
	jclient, _ := json.Marshal(cliente)
	f, err := os.OpenFile("data/clients.json", os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	f.WriteString(string(jclient))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	newClient(1, "Miguel", 1)
}
