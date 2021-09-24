package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/libsv/go-bt"
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

	//https://api.whatsonchain.com/v1/bsv/main/tx/fab0ba932003c3aaed63edc3db59256d8869d01cac1672cd1964b76c9fe29b4e/hex
	rawtx := "010000000252260e905b08cd928fda3ef55dd10bb2099c0d4dac414e4b063dfaffafd480ac0b0000006a47304402206c42d75930b0f5ce8dedf020509535c8be5b168757b95ba5a3566001069e5136022007be0b0bf73f1df12afaf3e291140755dbad3f7324dae5089d62d992055eec72412103107feff22788a1fc8357240bf450fd7bca4bd45d5f8bac63818c5a7b67b03876ffffffff5fb5cabd03bb6fd09a8ce20899b6e540a8c3ff34db067f20e68bf54f55bd0333020000006b483045022100ad94af60e699296eff611767ada0825382e4e1166810af4970f11f1ce0d30f3302204549a2841bc9585851cdc1dbb43a3e3aefe83c4b5cc094e2667c28e9b390d915412103c134c904118b148d32492cd17d1183088f708a3e4a7429f3260ff51b9e72c6ccffffffff030000000000000000fd0d01006a0372756e01050c63727970746f6669676874734cf67b22696e223a312c22726566223a5b22316531353761346133316431666565313566323665313666393861633931393436326164396439646134303563613566626433656431323233333237333862315f6f31225d2c226f7574223a5b2231656162323963313964353462383231373030356363653138386130376336343364353635303065336332623732376266383630326263653665623035633665225d2c2264656c223a5b5d2c22637265223a5b5d2c2265786563223a5b7b226f70223a2243414c4c222c2264617461223a5b7b22246a6967223a307d2c22626567696e222c5b313633323430353335313434355d5d7d5d7d11010000000000001976a9143b80a2d74a2b6dcd2f15fdea0d14aa58736de6d788ac00f50000000000001976a9144f7d6a485e09770f947c0ba38d15050a5a80b6fa88ac00000000"

	tx, err3 := bt.NewTxFromString(rawtx)
	if err3 != nil {
		log.Println(err3)
	}
	bufress, err5 := tx.MarshalJSON()
	log.Print(err5, string(bufress))
	// err4 := txa.AddOpReturnOutput(d)
	// if err4 != nil {
	// 	log.Print(err4)
	// }
	// log.Print(tx)
	// err4 := transaction.UnmarshalJSON(txJson)
	// if err4 != nil {
	// 	log.Print("Unable to unmarshall tx", err4)
	// }
	// log.Print(txJson)
}
