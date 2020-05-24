package discord

import "github.com/bwmarrin/discordgo"

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Players []*discordgo.User
	Civs    []*Civ
	Bans    map[*discordgo.User]*Civ
	Picks   map[*discordgo.User][]*Civ
}

// NewCivSession returns a clean CivSession.
func NewCivSession() *CivSession {
	return &CivSession{
		Players: []*discordgo.User{},
		Civs:    genCivs(),
		Bans:    make(map[*discordgo.User]*Civ, 0),
	}
}

// reset clears the CivSession referenced by the pointer receiver to the func.
func (cs *CivSession) reset() {
	cs.Players = []*discordgo.User{}
	cs.Bans = make(map[*discordgo.User]*Civ, 0)
}
