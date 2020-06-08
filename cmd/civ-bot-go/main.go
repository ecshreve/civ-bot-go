package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ecshreve/civ-bot-go/internal/bot"
)

func main() {
	token := os.Getenv("CIV_BOT_TOKEN")

	// Create a new Bot, exit if we encounter an error.
	b, err := bot.NewBot(token)
	if err != nil {
		log.Fatalf("error creating bot - %+v", err)
	}

	// Open the discordgo.Session for the Bot.
	err = b.StartSession()
	if err != nil {
		log.Fatalf("error starting bot session - %+v", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the Discord session for the Bot.
	b.EndSession()
}
