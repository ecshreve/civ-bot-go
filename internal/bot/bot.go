package bot

import (
	"github.com/samsarahq/go/oops"
)

// Bot holds data and implements functions for an instance of the civ-bot.
type Bot struct {
	*DiscordSession
	*Config
	*CivState
}

// NewBot takes a DiscordToken and returns a Bot.
func NewBot(token string) (*Bot, error) {
	ds, err := NewDiscordSession(token)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create DiscordSession")
	}

	config := NewConfig()

	// Initialize and return a new bot with command and reaction handlers.
	b := &Bot{
		DiscordSession: ds,
		Config:         config,
		CivState:       NewCivState(config),
	}
	b.Session.AddHandler(b.CommandHandler)
	b.Session.AddHandler(b.ReactionHandler)

	return b, nil
}

func (b *Bot) StartSession() error {
	err := b.Session.Open()
	if err != nil {
		return oops.Wrapf(err, "unable to open discord session")
	}

	return nil
}

func (b *Bot) EndSession() error {
	err := b.Session.Close()
	if err != nil {
		return oops.Wrapf(err, "unable to close discord session")
	}

	return nil
}
