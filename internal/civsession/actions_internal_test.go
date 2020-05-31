package civsession

import (
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/civ-bot-go/internal/civ"
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

func TestMakePick(t *testing.T) {
	clk := clock.NewMock()

	testcases := []struct {
		description       string
		civs              []constants.CivKey
		initialPickedCivs []constants.CivKey
		expectPick        bool
		expected          constants.CivKey
	}{
		{
			description: "empty input slice returns nil",
			civs:        []constants.CivKey{},
			expectPick:  false,
		},
		{
			description: "pick from slice of length 1 returns that item",
			civs:        []constants.CivKey{constants.AMERICA},
			expectPick:  true,
			expected:    constants.AMERICA,
		},
		{
			description:       "pick from slice of length 1 where that civ is already picked returns nil",
			civs:              []constants.CivKey{constants.AMERICA},
			initialPickedCivs: []constants.CivKey{constants.AMERICA},
			expectPick:        false,
		},
		// This testcase is deterministic because we provide a MockClock to our
		// CivSession, and we seed our random number generator in MakePick with
		// the current time.
		{
			description: "pick from slice of all civs returns indonesia",
			civs:        constants.CivKeys,
			expectPick:  true,
			expected:    constants.INDONESIA,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			cs := NewCivSession()
			cs.Clock = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, cs.CivMap[k])
			}

			for _, k := range testcase.initialPickedCivs {
				cs.CivMap[k].Picked = true
			}

			actual := cs.makePick(civsToTest)
			if testcase.expectPick {
				assert.Equal(t, cs.CivMap[testcase.expected], actual)
				assert.Same(t, cs.CivMap[testcase.expected], actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}
