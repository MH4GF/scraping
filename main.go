package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupEnv()
	result := scparing()
	monthlyTotalTable := newTotalTable(result.Eq(0).Text(), result.Eq(1).Text(), result.Eq(2).Text())
	client := newSlackClient(os.Getenv("SLACK_CHANNEL_ID"))
	client.postMessage(monthlyTotalTable)
}
