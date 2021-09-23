package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Utxo struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

func main() {
	p := Utxo{Txid: "3241545134251234", Vout: 5}
	pretty, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	log.Println("string", string(pretty))

	rawJSON, err2 := ioutil.ReadFile("utxo.json")
	if err2 != nil {
		log.Println(err2)
	}
	var m Utxo
	json.Unmarshal(rawJSON, &m)
	log.Println("marshall", m)
}
