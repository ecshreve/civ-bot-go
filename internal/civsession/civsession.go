package civsession

import (
	"fmt"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/bwmarrin/discordgo"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/ecshreve/civ-bot-go/internal/discord"
)

type civSession interface {
	banCiv(string, string) (*civ.Civ, error)
}

// CivConfig holds a config for a CivSession.
type CivConfig struct {
	NumBans        int
	NumPicks       int
	NumRepicks     int
	UseFilthyTiers bool
}

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Config           *CivConfig
	Clock            clock.Clock
	Players          []*discord.User
	PlayerMap        map[string]*discord.User
	Civs             []*civ.Civ
	CivMap           map[constants.CivKey]*civ.Civ
	Bans             map[string][]*civ.Civ
	Picks            map[string][]*civ.Civ
	PickTime         time.Time
	RePickVotes      int
	RePicksRemaining int
}

// DefaultCivConfig holds the default values for a CivConfig. This value is used
// when a new CivSession is created.
var DefaultCivConfig = CivConfig{
	NumBans:        1,
	NumPicks:       3,
	NumRepicks:     3,
	UseFilthyTiers: false,
}

// NewCivSession returns a new CivSession, note map fields are initialized to
// empty zero lengtrh maps.
func NewCivSession() *CivSession {
	cfg := DefaultCivConfig

	civs := civ.GenCivs()
	civMap := civ.GenCivMap(civs)

	return &CivSession{
		Config:           &cfg,
		Clock:            clock.New(),
		Players:          make([]*discord.User, 0),
		PlayerMap:        make(map[string]*discord.User),
		Civs:             civs,
		CivMap:           civMap,
		Bans:             make(map[string][]*civ.Civ),
		Picks:            make(map[string][]*civ.Civ),
		RePicksRemaining: cfg.NumRepicks,
	}
}

// Reset clears the CivSession referenced by the pointer receiver to the func,
// and maintains the CivConfig.
func (cs *CivSession) Reset() {
	civs := civ.GenCivs()
	civMap := civ.GenCivMap(civs)

	cs.Players = make([]*discord.User, 0)
	cs.PlayerMap = make(map[string]*discord.User)
	cs.Civs = civs
	cs.CivMap = civMap
	cs.Bans = make(map[string][]*civ.Civ)
	cs.Picks = make(map[string][]*civ.Civ)
	cs.PickTime = time.Time{}
	cs.RePickVotes = 0
	cs.RePicksRemaining = cs.Config.NumRepicks
}

// getConfigEmbedFields builds and returns the MessageEmbedFields for the
// CivSession's current CivConfig.
func (cs *CivSession) getConfigEmbedFields() []*discordgo.MessageEmbedField {
	maxPlayers := len(constants.CivKeys) / (cs.Config.NumBans + cs.Config.NumPicks)

	return []*discordgo.MessageEmbedField{
		{
			Name:  "`NumBans` -- the number of Civs each player gets to ban",
			Value: fmt.Sprintf("**%d**", cs.Config.NumBans),
		},
		{
			Name:  "`NumPicks` -- the number of Civs each player gets to choose from",
			Value: fmt.Sprintf("**%d**", cs.Config.NumPicks),
		},
		{
			Name:  "`NumRepicks` -- the max number of times allowed to re-pick Civs",
			Value: fmt.Sprintf("**%d**", cs.Config.NumRepicks),
		},
		{
			Name:  "`UseFilthyTiers` -- make picks based on FilthyRobot's tier list -- setting this to `true` ensures that each Player gets at minimum one t1/t2 Civ in their list of Picks",
			Value: fmt.Sprintf("**%v**", cs.Config.UseFilthyTiers),
		},
		{
			Name:  "Players",
			Value: fmt.Sprintf("This Config allows for a max of **%d** players", maxPlayers),
		},
	}
}
