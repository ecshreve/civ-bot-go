package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/stretchr/testify/assert"

	"github.com/samsarahq/go/snapshotter"
)

// fixupCivStateForSnapshot nils out the CivState Civs slice and CivMap map to
// make the Snapshots easier to validate.
func (b *Bot) fixupCivStateForSnapshot() {
	b.CivState.Civs = nil
	b.CivState.CivMap = nil
}

func TestNewState(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	output := NewCivState()
	snap.Snapshot("default civ state", output)
}

func TestReset(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, _ := MockBot(t)

	// Validate the default CivState and CivConfig.
	b.fixupCivStateForSnapshot()
	snap.Snapshot("new state", b.CivState)
	snap.Snapshot("new config", b.CivConfig)

	// Change some values in the CivState and CivConfig.
	b.CivState.RePickVotes = 99
	b.CivState.RePicksRemaining = 0
	b.CivConfig.Bans = 77
	b.CivConfig.Picks = 88
	b.CivConfig.RePicks = 99

	// Validate the altered CivState and CivConfig.
	b.fixupCivStateForSnapshot()
	snap.Snapshot("altered state", b.CivState)
	snap.Snapshot("altered config", b.CivConfig)

	// Reset maintaining the CivConfig.
	b.Reset(true)
	b.fixupCivStateForSnapshot()
	snap.Snapshot("state after reset - keepConfig true", b.CivState)
	snap.Snapshot("config after reset - keepConfig true", b.CivConfig)

	// Reset without maintaining the CivConfig.
	b.Reset(false)
	b.fixupCivStateForSnapshot()
	snap.Snapshot("state after reset - keepConfig false", b.CivState)
	snap.Snapshot("config after reset - keepConfig false", b.CivConfig)
}

func TestReadyToPick(t *testing.T) {
	b, _ := MockBot(t)
	b.CivConfig.Bans = 2

	testUserIDs := []string{"testPlayer1", "testPlayer2"}
	for _, id := range testUserIDs {
		testUser := &discordgo.User{ID: id}
		testPlayer := NewPlayer(testUser)
		b.CivState.Players = append(b.CivState.Players, testPlayer)
		b.CivState.PlayerMap[testPlayer.PlayerID] = testPlayer
	}

	testcases := []struct {
		description  string
		existingBans map[PlayerID][]constants.CivKey
		expected     bool
	}{
		{
			description: "no bans",
			expected:    false,
		},
		{
			description: "one ban for one player",
			existingBans: map[PlayerID][]constants.CivKey{
				b.CivState.Players[0].PlayerID: []constants.CivKey{constants.AMERICA},
			},
			expected: false,
		},
		{
			description: "one ban for each player",
			existingBans: map[PlayerID][]constants.CivKey{
				b.CivState.Players[0].PlayerID: []constants.CivKey{constants.AMERICA},
				b.CivState.Players[1].PlayerID: []constants.CivKey{constants.BRAZIL},
			},
			expected: false,
		},
		{
			description: "two bans for one player and one ban for the other player",
			existingBans: map[PlayerID][]constants.CivKey{
				b.CivState.Players[0].PlayerID: []constants.CivKey{constants.AMERICA, constants.ZULUS},
				b.CivState.Players[1].PlayerID: []constants.CivKey{constants.BRAZIL},
			},
			expected: false,
		},
		{
			description: "two bans for each player",
			existingBans: map[PlayerID][]constants.CivKey{
				b.CivState.Players[0].PlayerID: []constants.CivKey{constants.AMERICA, constants.ZULUS},
				b.CivState.Players[1].PlayerID: []constants.CivKey{constants.BRAZIL, constants.KOREA},
			},
			expected: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			banMap := make(map[PlayerID][]*civ.Civ)
			for k, v := range testcase.existingBans {
				var playerBans []*civ.Civ
				for _, ck := range v {
					playerBans = append(playerBans, b.CivState.CivMap[ck])
				}
				banMap[k] = playerBans
			}
			b.CivState.Bans = banMap

			actual := b.ReadyToPick()
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
