package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func renderOutputJson() {
	result := scparing()
	monthlyTotalTable := newTotalTable(result.Eq(0).Text(), result.Eq(1).Text(), result.Eq(2).Text())
	outputJson, err := json.Marshal(&monthlyTotalTable)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(outputJson))
}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupEnv()
	renderOutputJson()
}
