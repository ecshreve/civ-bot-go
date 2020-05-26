package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var Session = NewCivSession()

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Players     map[string]*discordgo.User
	Civs        []*Civ
	Bans        map[string]*Civ
	Picks       map[*discordgo.User][]*Civ
	PickTime    time.Time
	RePickVotes int
}

// NewCivSession returns a new CivSession, note map fields are initialized to
// empty zero lengtrh maps.
func NewCivSession() *CivSession {
	return &CivSession{
		Players: make(map[string]*discordgo.User),
		Civs:    genCivs(),
		Bans:    make(map[string]*Civ),
		Picks:   make(map[*discordgo.User][]*Civ),
	}
}

// reset clears the CivSession referenced by the pointer receiver to the func.
func (cs *CivSession) reset() {
	cs.Players = make(map[string]*discordgo.User)
	cs.Civs = genCivs()
	cs.Bans = make(map[string]*Civ)
	cs.Picks = make(map[*discordgo.User][]*Civ)
	cs.PickTime = time.Time{}
	cs.RePickVotes = 0
}

// banCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func banCiv(civToBan string, userID string) *Civ {
	cs := Session
	c := cs.getCivByString(civToBan)
	if c == nil || c.Banned == true {
		return nil
	}

	// If this player had previously banned a Civ then unban the previous Civ.
	if _, ok := cs.Bans[userID]; ok {
		cs.Bans[userID].Banned = false
	}

	c.Banned = true
	cs.Bans[userID] = c

	return c
}
