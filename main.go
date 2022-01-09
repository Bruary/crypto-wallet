package main

import (
	"fmt"

	"github.com/Bruary/crypto-wallet/db"
	"github.com/Bruary/crypto-wallet/models"
	"github.com/Bruary/crypto-wallet/service/users"
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

	svc := users.NewUsers()

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

	user.Post("/create", func(c *fiber.Ctx) error {

		c.Context().SetContentType("application/json")

		req := models.User{}
		if err := UnmarshalRequest(c, &req); err != nil {
			return err
		}

		resp := svc.CreateUser(req)

		err = MarshalAndSetResponseBody(c, resp)
		if err != nil {
			return err
		}

		return nil
	})

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

func UnmarshalRequest(c *fiber.Ctx, req interface{}) error {

	if err := json.Unmarshal(c.Body(), req); err != nil {
		c.SendString("Error while unmarshaling.")
		return err
	}

	return nil
}

func MarshalAndSetResponseBody(c *fiber.Ctx, resp interface{}) error {

	result, err := json.Marshal(resp)
	if err != nil {
		c.SendString("Marshaling failed.")
		return err
	}

	c.Context().SetBody(result)
	return nil
}
