package discord

import "github.com/bwmarrin/discordgo"

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Players []*discordgo.User
	Civs    []*Civ
	Picks   map[*discordgo.User][]*Civ
}

// NewCivSession returns a clean CivSession.
func NewCivSession() *CivSession {
	return &CivSession{
		Players: []*discordgo.User{},
		Civs:    genCivs(),
	}
}

// reset clears the CivSession referenced by the pointer receiver to the func.
func (cs *CivSession) reset() {
	cs.Players = []*discordgo.User{}
}
