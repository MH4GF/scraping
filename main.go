package main

import (
	"github.com/joho/godotenv"
	"log"
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
	monthlyTotalTable.renderingJson()
}
