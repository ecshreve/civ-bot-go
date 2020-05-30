package civsession

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/civ-bot-go/internal/constants"
)

func TestBanCiv(t *testing.T) {
	testdata := NewTestData()
	players := testdata.Players
	playerMap := make(map[string]*discordgo.User)
	for _, p := range players {
		playerMap[p.ID] = p
	}

	testcases := []struct {
		description string
		civToBan    string
		expectBan   bool
		expected    constants.CivKey
	}{
		{
			description: "empty string expect nil",
			civToBan:    "",
			expectBan:   false,
		},
		{
			description: "ban by civ name",
			civToBan:    "america",
			expectBan:   true,
			expected:    constants.AMERICA,
		},
		{
			description: "ban by leader name",
			civToBan:    "washington",
			expectBan:   true,
			expected:    constants.AMERICA,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			cs := NewCivSession()
			cs.Players = playerMap
			testPlayer := players[0]
			expectedCiv := cs.CivMap[testcase.expected]

			actual := cs.banCiv(testcase.civToBan, testPlayer.ID)
			if testcase.expectBan {
				// Make sure we only banned the expected Civ and all others are
				// not banned.
				for _, c := range cs.Civs {
					if c.Key == expectedCiv.Key {
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
					if c.Key == expectedCiv.Key {
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
