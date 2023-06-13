package main

import (
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	event_consumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

// it is better to get host just as we get Token - through the flag
// but for lack of time lets skip it for a while

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
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
