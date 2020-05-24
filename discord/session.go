package discord

import "github.com/bwmarrin/discordgo"

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Players map[string]*discordgo.User
	Civs    []*Civ
	Bans    map[string]*Civ
	Picks   map[*discordgo.User][]*Civ
}

// NewCivSession returns a new CivSession, note map fields are initialized to
// empty zero lengtrh maps.
func NewCivSession() *CivSession {
	return &CivSession{
		Players: map[string]*discordgo.User{},
		Civs:    genCivs(),
		Bans:    map[string]*Civ{},
	}
}

// reset clears the CivSession referenced by the pointer receiver to the func.
func (cs *CivSession) reset() {
	cs.Players = map[string]*discordgo.User{}
	cs.Bans = make(map[string]*Civ, 0)
}
