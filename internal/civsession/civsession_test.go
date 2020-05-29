package civsession_test

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/internal/civsession"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestNewCivSession(t *testing.T) {
	cs := civsession.NewCivSession()
	assert.Equal(t, 0, len(cs.Players))
	assert.Equal(t, 43, len(cs.Civs))
	assert.Equal(t, 0, len(cs.Bans))
	assert.Equal(t, 0, len(cs.Picks))
	assert.True(t, cs.PickTime.IsZero())
	assert.Equal(t, 0, cs.RePickVotes)
	assert.Equal(t, 3, cs.RePicksRemaining)
}

func TestReset(t *testing.T) {
	testdata := civsession.NewTestData()

	testcases := []struct {
		description string
		cs          *civsession.CivSession
	}{
		{
			description: "reset a new CivSession",
			cs:          testdata.CS,
		},
		{
			description: "reset a CivSession that has Players",
			cs:          testdata.CSWithPlayers,
		},
		{
			description: "reset a CivSession that has Players and Bans",
			cs:          testdata.CSWithPlayersAndBans,
		},
		{
			description: "reset a CivSession that has Players, Bans, and Picks, a PickTime, and RePickVotes",
			cs:          testdata.CSWithPlayersBansAndPicks,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := testcase.cs
			actual.Reset()

			assert.Equal(t, 0, len(actual.Players))
			assert.Equal(t, len(constants.CivKeys), len(actual.Civs))

			for _, civ := range actual.Civs {
				assert.False(t, civ.Banned)
				assert.False(t, civ.Picked)
			}

			assert.Equal(t, 0, len(actual.Bans))
			assert.Equal(t, 0, len(actual.Picks))
			assert.True(t, actual.PickTime.IsZero())
			assert.Equal(t, 0, actual.RePickVotes)
			assert.Equal(t, 3, actual.RePicksRemaining)
		})
	}
}
