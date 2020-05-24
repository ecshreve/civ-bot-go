package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ecshreve/civ-bot-go/discord"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("CIV_BOT_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Create a CivSession and register handlers.
	cs := discord.NewCivSession()
	dg.AddHandler(cs.CommandsHandler)
	dg.AddHandler(cs.ReactionsHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
