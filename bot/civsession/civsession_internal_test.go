package civsession

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/bot/constants"
	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	testdata := NewTestData()

	testcases := []struct {
		description string
		cs          *CivSession
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
			cs:          testdata.csWithPlayersBansAndPicks,
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
		})
	}
}
