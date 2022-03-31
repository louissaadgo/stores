package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"stores/models"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStoreSignup(t *testing.T) {

	go func() {
		app := fiber.New()
		app.Post("/api/auth/store/signup", StoreSignup)
		app.Listen(":4000")
	}()

	correctData := models.StoreSignup{
		Name:        "Store 1",
		Address:     "Address 1",
		CountryCode: "961",
		Phone:       "1111111",
		PublicEmail: "pubemail@test.com",
		Email:       "email@test.com",
		Password:    "password",
	}
	url := "http://localhost:4000/api/auth/store/signup"
	jsonData, err := json.Marshal(correctData)
	if err != nil {
		t.Errorf("Failed converting to json")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected res %s", resp.Status)
	}
}
