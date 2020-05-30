package civsession

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/internal/constants"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"

	"github.com/stretchr/testify/assert"
)

func TestBanCiv(t *testing.T) {
	testdata := NewTestData()
	players := testdata.Players
	playerMap := make(map[string]*discordgo.User)
	for _, p := range players {
		playerMap[p.ID] = p
	}
	civMap := civ.GenCivMap()

	testcases := []struct {
		description string
		civToBan    string
		expected    *civ.Civ
	}{
		{
			description: "empty string expect nil",
			civToBan:    "",
			expected:    nil,
		},
		{
			description: "ban by civ name",
			civToBan:    "america",
			expected:    civMap[constants.AMERICA],
		},
		{
			description: "ban by leader name",
			civToBan:    "washington",
			expected:    civMap[constants.AMERICA],
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			cs := NewCivSession()
			cs.Players = playerMap
			testPlayer := players[0]

			actual := cs.banCiv(testcase.civToBan, testPlayer.ID)
			if testcase.expected != nil {
				// Make sure we only banned the expected Civ and all others are
				// not banned.
				for _, c := range cs.Civs {
					if c.Key == testcase.expected.Key {
						assert.True(t, c.Banned)
					} else {
						assert.False(t, c.Banned)
					}
				}

				// Make sure our expected Civ is the only item in the CivSession
				// Ban slice.
				assert.Equal(t, 1, len(cs.Bans))
				foundInBans := false
				for _, c := range cs.Bans[testPlayer.ID] {
					if c.Key == testcase.expected.Key {
						foundInBans = true
						break
					}
				}
				assert.True(t, foundInBans)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}
