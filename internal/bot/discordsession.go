package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"
)

type DiscordSession struct {
	discordgo.Session
}

func NewDiscordSession(token string) (*DiscordSession, error) {
	// Create a new Discord session using the provided token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create new discordgo.Session")
	}

	ds := DiscordSession{*dg}
	return &ds, nil
}
