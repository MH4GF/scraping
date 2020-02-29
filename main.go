package main

import (
	"os"
)

func main() {
	result := scparing()
	monthlyTotalTable := newTotalTable(result.Eq(0).Text(), result.Eq(1).Text(), result.Eq(2).Text())
	client := newSlackClient(os.Getenv("SLACK_CHANNEL_ID"))
	client.postMessage(monthlyTotalTable)
}
