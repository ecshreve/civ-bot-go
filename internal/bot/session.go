package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"
)

type DiscordSession struct {
	discordgo.Session
}

func NewDiscordSession(token string) (*DiscordSession, error) {
	// Create a new Discord session using the provided bot token, if we
	// encounter an error log it and exit.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create new discordgo.Session")
	}

	ds := &DiscordSession{*dg}

	return ds, nil
}

// Open opens a websocket connection to Discord and begins
// listening, if we encounter an error then we return it to the caller.
func (s *DiscordSession) Open() error {
	err := s.Session.Open()
	if err != nil {
		return oops.Wrapf(err, "unable to open discordgo.Session")
	}

	return nil
}

// Close cleanly closes down the Discord session.
func (s *DiscordSession) Close() error {
	err := s.Session.Close()
	if err != nil {
		log.Printf("unable do close discordgo.Session: %+v\n", err)
		return oops.Wrapf(err, "unable do close discordgo.Session")
	}

	return nil
}

// AddHandler attaches a handler function to the DiscordSession.
func (s *DiscordSession) AddHandler(handler interface{}) func() {
	return s.Session.AddHandler(handler)
}
