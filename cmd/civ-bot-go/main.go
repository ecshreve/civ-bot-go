package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ecshreve/civ-bot-go/internal/civsession"
	"github.com/ecshreve/civ-bot-go/internal/discord"
)

func main() {
	token := os.Getenv("CIV_BOT_TOKEN")

	// Create a new Discord session using the provided bot token, if we
	// encounter an error log it and exit.
	dg, err := discord.NewSessDAL("Bot " + token)
	if err != nil {
		log.Fatalf("error creating Discord session - %+v", err)
	}

	// Create a new CivSession, and attach handlers to the DiscordSession.
	cs := civsession.NewCivSession()
	dg.AddHandler(cs.CommandsHandler)
	dg.AddHandler(cs.ReactionsHandler)

	// Open a websocket connection to Discord and begin listening, if we
	// encounter an error log it and exit.
	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening connection - %+v", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
