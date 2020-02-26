package main

import (
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"log"
	"os"
)

type Endpoints struct {
	client            *slack.Client
	verificationToken string
}

type SlackBotService interface {
	create() *slack.Client
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func create() *Endpoints {
	verificationToken := goDotEnvVariable("SLACK_VERIFICATION_TOKEN")
	client := slack.New(goDotEnvVariable("SLACK_API"))
	return &Endpoints{client, verificationToken}
}
