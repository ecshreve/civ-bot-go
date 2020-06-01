package civsession_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/civ-bot-go/internal/civsession"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

func TestNewCivSession(t *testing.T) {
	cs := civsession.NewCivSession()

	// Assert that the newly created CivSession config has the same values as
	// the DefaultCivConfig, but that they don't point to the same value.
	assert.EqualValues(t, civsession.DefaultCivConfig, *cs.Config)
	assert.NotSame(t, &civsession.DefaultCivConfig, cs.Config)

	// Assert that a new Clock was created for the CivSession.
	assert.NotNil(t, cs.Clock)

	// Assert that the slices and maps were initialized.
	assert.Equal(t, 0, len(cs.Players))
	assert.Equal(t, 0, len(cs.PlayerMap))
	assert.Equal(t, len(constants.CivKeys), len(cs.Civs))
	assert.Equal(t, len(constants.CivKeys), len(cs.CivMap))
	assert.Equal(t, 0, len(cs.Bans))
	assert.Equal(t, 0, len(cs.Picks))

	// Assert the CivSession variables were initialized to the correct values.
	assert.True(t, cs.PickTime.IsZero())
	assert.Equal(t, 0, cs.RePickVotes)
	assert.Equal(t, civsession.DefaultCivConfig.NumRepicks, cs.RePicksRemaining)
}

func TestReset(t *testing.T) {
	testdata := civsession.NewTestData()

	testcases := []struct {
		description string
		cs          *civsession.CivSession
		config      *civsession.CivConfig
	}{
		{
			description: "reset a new CivSession",
			cs:          testdata.CS,
		},
		{
			description: "reset a new CivSession after changing the config",
			cs:          testdata.CS,
			config: &civsession.CivConfig{
				NumBans:        111,
				NumPicks:       111,
				NumRepicks:     111,
				UseFilthyTiers: true,
			},
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

			// Set the CivConfig.
			expectedCivConfig := &civsession.DefaultCivConfig
			if testcase.config != nil {
				actual.Config.NumBans = testcase.config.NumBans
				actual.Config.NumPicks = testcase.config.NumPicks
				actual.Config.NumRepicks = testcase.config.NumRepicks
				actual.Config.UseFilthyTiers = testcase.config.UseFilthyTiers
				expectedCivConfig = testcase.config
			}

			// Reset the CivSession.
			actual.Reset()

			// Assert that the CivSession config has the expected values, we
			// don't expect the CivConfig to change after a reset.
			assert.EqualValues(t, expectedCivConfig, actual.Config)
			assert.NotSame(t, expectedCivConfig, actual.Config)

			// Assert that the slices and maps were re-initialized correctly.
			assert.Equal(t, 0, len(actual.Players))
			assert.Equal(t, 0, len(actual.PlayerMap))
			assert.Equal(t, len(constants.CivKeys), len(actual.Civs))
			assert.Equal(t, len(constants.CivKeys), len(actual.CivMap))
			assert.Equal(t, 0, len(actual.Bans))
			assert.Equal(t, 0, len(actual.Picks))

			// Assert that the CivSession Civs were reset correctly.
			for _, civ := range actual.Civs {
				assert.False(t, civ.Banned)
				assert.False(t, civ.Picked)
			}

			// Assert the CivSession variables were re-initialized to the
			// correct vals.
			assert.True(t, actual.PickTime.IsZero())
			assert.Equal(t, 0, actual.RePickVotes)
			assert.Equal(t, expectedCivConfig.NumRepicks, actual.RePicksRemaining)
		})
	}
}
