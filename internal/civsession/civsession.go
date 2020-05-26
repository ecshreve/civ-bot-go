package civsession

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
)

// CS stores the global CivSession.
var CS = NewCivSession()

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Players     map[string]*discordgo.User
	Civs        []*civ.Civ
	Bans        map[string]*civ.Civ
	Picks       map[*discordgo.User][]*civ.Civ
	PickTime    time.Time
	RePickVotes int
}

// NewCivSession returns a new CivSession, note map fields are initialized to
// empty zero lengtrh maps.
func NewCivSession() *CivSession {
	return &CivSession{
		Players: make(map[string]*discordgo.User),
		Civs:    civ.GenCivs(),
		Bans:    make(map[string]*civ.Civ),
		Picks:   make(map[*discordgo.User][]*civ.Civ),
	}
}

// Reset clears the CivSession referenced by the pointer receiver to the func.
func (cs *CivSession) Reset() {
	cs.Players = make(map[string]*discordgo.User)
	cs.Civs = civ.GenCivs()
	cs.Bans = make(map[string]*civ.Civ)
	cs.Picks = make(map[*discordgo.User][]*civ.Civ)
	cs.PickTime = time.Time{}
	cs.RePickVotes = 0
}

// BanCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func BanCiv(civToBan string, userID string) *civ.Civ {
	cs := CS
	c := civ.GetCivByString(civToBan, cs.Civs)
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
