package main

import (
	"fmt"

	"github.com/Bruary/crypto-wallet/db"
	"github.com/Bruary/crypto-wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// TODO
// Create structs for users, accounts, etc..
// Add table to database
// Create handlers

func main() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Loading .env failed" + err.Error())
	}

	// Connect to db
	db.ConnectToDB()

	app := fiber.New()

	v1 := app.Group("/v1")

	user := v1.Group("/user")
	user.Post("/create")
	user.Post("/deactivate")

	investment := v1.Group("/investment")
	investment.Post("")

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
