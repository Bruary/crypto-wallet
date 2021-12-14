package main

import (
	"fmt"

	"github.com/Bruary/crypto-wallet/models"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// main functionalities
//	1) Know if you are in profit or loss.
//  2) calculate the price of currency that when you will start gaining profit.
//  3) Add the percentage of profit so you get notified.

func main() {

	currencies, err := GetCurrencies()
	if err != nil {
		log.Fatalln(err)
	}

	respJson, err := json.MarshalIndent(currencies, "", "    ")

	fmt.Println(string(respJson))

}

func GetCurrencies() (interface{}, error) {

	// Make the Get method call to Nomics
	resp, err := http.Get("https://api.nomics.com/v1/currencies/ticker?key=m_875081ad0011453842df0b214e161109072444d6&ids=ETH,BTC,XPR,LTC&interval=1h,1d,7d,30d,365d&convert=USD&per-page=100&page=1")
	if err != nil {
		return nil, err
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var nomicsResponse models.Currency

	err = json.Unmarshal(body, &nomicsResponse)
	if err != nil {
		return nil, err
	}

	return nomicsResponse, nil
}
