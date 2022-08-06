package main

import (
	"fmt"
	"log"
	"os"
    "github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})


	message := "Ola, este SMS foi enviado com Go , e , com Api Twilio "

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TWILIO_TO"))
	params.SetFrom(os.Getenv("TWILIO_FROM")) 
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println("Message Sid: " + *resp.Sid)
	}
}