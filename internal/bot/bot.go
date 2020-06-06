package bot

import (
	"github.com/samsarahq/go/oops"
)

// Bot holds data and implements functions for an instance of the civ-bot.
type Bot struct {
	DS       *DiscordSession
	Config   *Config
	CivState *CivState
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
		DS:       ds,
		Config:   config,
		CivState: NewCivState(config),
	}

	return b, nil
}

func (b *Bot) AddHandlers() error {
	if b.DS == nil {
		return oops.Errorf("unable to add handlers to nil discordgo.Session")
	}

	b.DS.AddHandler(b.CommandHandler)
	b.DS.AddHandler(b.ReactionHandler)

	return nil
}

func (b *Bot) StartSession() error {
	err := b.DS.Open()
	if err != nil {
		return oops.Wrapf(err, "unable to open discord session")
	}

	return nil
}

func (b *Bot) EndSession() error {
	err := b.DS.Close()
	if err != nil {
		return oops.Wrapf(err, "unable to close discord session")
	}

	return nil
}
