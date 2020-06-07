package bot

import (
	"github.com/samsarahq/go/oops"
)

// Bot holds data and implements functions for an instance of the civ-bot.
type Bot struct {
	DS         *DiscordSession
	Commands   []Command
	CommandMap map[CommandID]Command
	CivConfig  *CivConfig
	CivState   *CivState
}

// NewBot takes a token and returns a Bot.
func NewBot(token string) (*Bot, error) {
	// Create a new DiscordSession for the bot, return if we encounter an error.
	ds, err := NewDiscordSession(token)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create DiscordSession")
	}

	// Initialize fields for a new Bot.
	b := &Bot{
		DS:        ds,
		CivConfig: NewCivConfig(),
		CivState:  NewCivState(),
		Commands:  AllCommands,
	}
	b.CommandMap = getCommandIDToCommandMap(b.Commands)

	// Attach message and reaction handlers to the bot.
	err = b.AddHandlers()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to add handlers to Bot")
	}

	return b, nil
}

// AddHandlers attaches handler functions to the Bot's DiscordSession.
func (b *Bot) AddHandlers() error {
	if b.DS == nil {
		return oops.Errorf("unable to add handlers to nil discordgo.Session")
	}

	b.DS.AddHandler(b.MessageHandler)
	b.DS.AddHandler(b.ReactionHandler)

	return nil
}

// StartSession is a wrapper around the DiscordSession Open() func.
func (b *Bot) StartSession() error {
	err := b.DS.Open()
	if err != nil {
		return oops.Wrapf(err, "unable to open discord session")
	}

	return nil
}

// EndSession is a wrapper around the DisccordSession Close() func.
func (b *Bot) EndSession() error {
	err := b.DS.Close()
	if err != nil {
		return oops.Wrapf(err, "unable to close discord session")
	}

	return nil
}
