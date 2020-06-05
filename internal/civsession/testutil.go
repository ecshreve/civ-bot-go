package civsession

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/discord"
)

func genPlayersForTest() []*discord.User {
	var players []*discord.User
	for i := 1; i <= 10; i++ {
		p := &discord.User{&discordgo.User{
			ID:       fmt.Sprintf("p%dID", i),
			Email:    fmt.Sprintf("p%dEmail@devnull.com", i),
			Username: fmt.Sprintf("p%dUsername", i),
		}}
		players = append(players, p)
	}
	return players
}

// setCivSessionPlayersForTest sets the CivSession's Players field based on the
// given slice of players.
func (cs *CivSession) setCivSessionPlayersForTest(players []*discord.User) {
	playerMap := make(map[string]*discord.User)
	for _, p := range players {
		playerMap[p.ID] = p
	}
	cs.PlayerMap = playerMap
	cs.Players = players
}

// setCivSessionBansForTest sets the CivSession's Bans field based on the
// given slice of players, and updates the CivSession's slice of Civs accordingly.
func (cs *CivSession) setCivSessionBansForTest(players []*discord.User) {
	for i, p := range players {
		cs.Bans[p.ID] = append(cs.Bans[p.ID], cs.Civs[i])
		cs.Civs[i].Banned = true
	}
}

// TODO error checking / edge case handling.
//
// setCivSessionPicksForTest sets the CivSession's Picks field based on the
// given slice of players, and updates the CivSession's slice of Civs accordingly.
// It also sets the PickTime and RePickVotes fields.
func (cs *CivSession) setCivSessionPicksForTest(players []*discord.User) {
	civInd := 0
	for _, p := range players {
		var picks []*civ.Civ
		for len(picks) < 3 && civInd < len(cs.Civs) {
			curCiv := cs.Civs[civInd]
			if curCiv.Banned == false && curCiv.Picked == false {
				curCiv.Picked = true
				picks = append(picks, curCiv)
			}
			civInd++
		}
		cs.Picks[p.ID] = picks
	}
	cs.PickTime = time.Now()
	cs.RePickVotes = 6
}

// CivBotTestData contains data to use in automated tests.
type CivBotTestData struct {
	Players                   []*discord.User
	CS                        *CivSession
	CSWithPlayers             *CivSession
	CSWithPlayersAndBans      *CivSession
	CSWithPlayersBansAndPicks *CivSession
}

// NewTestData returns a fresh instance of CivBotTestData to use in tests.
func NewTestData() *CivBotTestData {
	players := genPlayersForTest()

	csWithPlayers := NewCivSession()
	csWithPlayers.setCivSessionPlayersForTest(players)

	csWithPlayersAndBans := NewCivSession()
	csWithPlayersAndBans.setCivSessionPlayersForTest(players)
	csWithPlayersAndBans.setCivSessionBansForTest(players)

	csWithPlayersBansAndPicks := NewCivSession()
	csWithPlayersBansAndPicks.setCivSessionPlayersForTest(players)
	csWithPlayersBansAndPicks.setCivSessionBansForTest(players)
	csWithPlayersBansAndPicks.setCivSessionPicksForTest(players)

	data := &CivBotTestData{
		Players:                   players,
		CS:                        NewCivSession(),
		CSWithPlayers:             csWithPlayers,
		CSWithPlayersAndBans:      csWithPlayersAndBans,
		CSWithPlayersBansAndPicks: csWithPlayersBansAndPicks,
	}

	return data
}
