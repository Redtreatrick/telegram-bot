package main

import (
	"context"
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	event_consumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/sqlite"
)

// it is better to get host just as we get Token - through the flag
// but for lack of time lets skip it for a while

const (
	tgBotHost         = "api.telegram.org"
	storagePath       = "files_storage"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	//s := files.New(storagePath),
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can't connect ot storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatalf("can't init storage: ", err)
	}
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
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
		"tg-bot-token",
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
