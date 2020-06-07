package bot

import (
	"github.com/samsarahq/go/oops"
)

// Bot holds data and implements functions for an instance of the civ-bot.
type Bot struct {
	DS        *DiscordSession
	Commands  map[CommandID]Command
	CivConfig *CivConfig
	CivState  *CivState
}

// NewBot takes a token and returns a Bot.
func NewBot(token string) (*Bot, error) {
	// Create a new DiscordSession for the bot, return if we encounter an error.
	ds, err := NewDiscordSession(token)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create DiscordSession")
	}

	// Initialize and return a new Bot.
	b := &Bot{
		DS:        ds,
		CivConfig: NewCivConfig(),
		CivState:  NewCivState(),
		Commands: map[CommandID]Command{
			CommandID("help"): &helpCommand{},
		},
	}

	// Attach command and reaction handlers to the bot.
	err = b.AddHandlers()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to add handlers to Bot")
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
