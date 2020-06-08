package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"
)

// DiscordSession overrides the discordgo.Session, allowing us to extend it and
// mock the base class methods for testing.
type DiscordSession struct {
	*discordgo.Session
}

// NewDiscordSession returns a DiscordSession created with the given token.
func NewDiscordSession(token string) (*DiscordSession, error) {
	// Create a new Discord session using the provided token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create new discordgo.Session")
	}

	ds := DiscordSession{dg}
	return &ds, nil
}
