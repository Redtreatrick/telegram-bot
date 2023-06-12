package main

import (
	"flag"
	"log"
	"read-adviser-bot/clients/telegram"
)

// it is better to get host just as we get Token - through the flag
// but for lack of time lets skip it for a while

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	_ = tgClient

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

// name "must" implies that function will go panic instead of returning error
func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	// it is usually dangerous to compare pointer to smth
	// if it points to no element - the panic happens
	// but in this case its fine
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
