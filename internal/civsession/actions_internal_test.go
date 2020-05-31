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
		description       string
		civToBan          string
		playerID          string
		initialBannedCivs []constants.CivKey
		expectError       bool
		expected          constants.CivKey
	}{
		{
			description: "empty civToBan expect error",
			civToBan:    "",
			playerID:    players[0].ID,
			expectError: true,
		},
		{
			description: "empty userID expect error",
			civToBan:    "america",
			playerID:    "",
			expectError: true,
		},
		{
			description:       "civ already banned expect error",
			civToBan:          "america",
			playerID:          players[0].ID,
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			expectError:       true,
		},
		{
			description: "ban by civ name",
			civToBan:    "america",
			playerID:    players[0].ID,
			expected:    constants.AMERICA,
		},
		{
			description: "ban by leader name",
			civToBan:    "washington",
			playerID:    players[0].ID,
			expected:    constants.AMERICA,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			cs := NewCivSession()
			cs.Players = players
			cs.PlayerMap = playerMap

			for _, b := range testcase.initialBannedCivs {
				cs.CivMap[b].Banned = true
			}

			expectedCiv := cs.CivMap[testcase.expected]
			actual, err := cs.banCiv(testcase.civToBan, testcase.playerID)

			if !testcase.expectError {
				assert.NoError(t, err)
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
				for _, c := range cs.Bans[testcase.playerID] {
					if c.Key == expectedCiv.Key {
						foundInBans = true
						break
					}
				}
				assert.True(t, foundInBans)
			} else {
				assert.Nil(t, actual)
				assert.Error(t, err)
			}
		})
	}
}

func TestMakePick(t *testing.T) {
	clk := clock.NewMock()

	testcases := []struct {
		description       string
		civs              []constants.CivKey
		initialBannedCivs []constants.CivKey
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
		{
			description:       "pick from slice of length 1 where that civ is banned returns nil",
			civs:              []constants.CivKey{constants.AMERICA},
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			expectPick:        false,
		},
		{
			description:       "pick from slice of length 2 where one civ is banned, one is picked, return nil",
			civs:              []constants.CivKey{constants.AMERICA, constants.ARABIA},
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			initialPickedCivs: []constants.CivKey{constants.ARABIA},
			expectPick:        false,
		},
		{
			description:       "pick from slice of length 2 where one civ is banned returns the other civ",
			civs:              []constants.CivKey{constants.AMERICA, constants.ARABIA},
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			expectPick:        true,
			expected:          constants.ARABIA,
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

			for _, k := range testcase.initialBannedCivs {
				cs.CivMap[k].Banned = true
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

func TestMakePicks(t *testing.T) {
	clk := clock.NewMock()

	testcases := []struct {
		description       string
		civs              []constants.CivKey
		numPicks          int
		initialPickedCivs []constants.CivKey
		expected          []constants.CivKey
	}{
		{
			description: "empty input slice returns nil",
			civs:        []constants.CivKey{},
			numPicks:    1,
		},
		{
			description: "input slice less than numPicks returns nil",
			civs:        []constants.CivKey{constants.AMERICA},
			numPicks:    2,
		},
		{
			description:       "picking 2 civs from list of 2 with one already picked returns nil",
			civs:              []constants.CivKey{constants.AMERICA, constants.ARABIA},
			numPicks:          2,
			initialPickedCivs: []constants.CivKey{constants.AMERICA},
		},
		{
			description: "input slice with numPicks items returns those items",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA},
			numPicks:    2,
			expected:    []constants.CivKey{constants.AMERICA, constants.ARABIA},
		},
		// This testcase is deterministic because we provide a MockClock to our
		// CivSession, and we seed our random number generator in MakePick with
		// the current time.
		{
			description: "picking three civs from the full list returns 3 expected civs",
			civs:        constants.CivKeys,
			numPicks:    3,
			expected:    []constants.CivKey{constants.INDONESIA, constants.IROQUOIS, constants.MAYANS},
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

			var expectedCivs []*civ.Civ
			for _, k := range testcase.expected {
				expectedCivs = append(expectedCivs, cs.CivMap[k])
			}

			actual := cs.makePicks(civsToTest, testcase.numPicks)
			assert.EqualValues(t, expectedCivs, actual)
		})
	}
}

func TestMakePicksWithTier(t *testing.T) {
	testdata := NewTestData()
	players := testdata.Players
	clk := clock.NewMock()

	testcases := []struct {
		description string
		civs        []constants.CivKey
		players     []*discordgo.User
		numPicks    int
		expected    map[string][]constants.CivKey
	}{
		{
			description: "nil civ list returns nil",
			civs:        nil,
			players:     players[:1],
			numPicks:    1,
			expected:    nil,
		},
		{
			description: "one pick, one player, one low tier civ returns nil",
			civs:        []constants.CivKey{constants.AMERICA},
			players:     players[:1],
			numPicks:    1,
			expected:    nil,
		},
		{
			description: "one pick, one player, one top tier civ returns that civ",
			civs:        []constants.CivKey{constants.ARABIA},
			players:     players[:1],
			numPicks:    1,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.ARABIA},
			},
		},
		{
			description: "two picks, one player, one top tier civ, one low tier civ returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA},
			players:     players[:1],
			numPicks:    2,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.ARABIA, constants.AMERICA},
			},
		},
		{
			description: "one pick, two players, two top tier civ, two low tier civ returns the top tier civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     players[:2],
			numPicks:    1,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.AZTECS},
				players[1].ID: {constants.ARABIA},
			},
		},
		{
			description: "two picks, two players, two top tier civ, two low tier civ returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     players[:2],
			numPicks:    2,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.AZTECS, constants.AMERICA},
				players[1].ID: {constants.ARABIA, constants.ASSYRIA},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			// TODO: all of this setup can probably be pulled into a helper.
			cs := NewCivSession()
			cs.Clock = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, cs.CivMap[k])
			}
			civMap := civ.GenCivMap(civsToTest)
			cs.Civs = civsToTest
			cs.CivMap = civMap

			var playerMap = make(map[string]*discordgo.User)
			for _, u := range testcase.players {
				playerMap[u.ID] = u
			}
			cs.Players = testcase.players
			cs.PlayerMap = playerMap

			cs.Config.NumPicks = testcase.numPicks
			cs.Config.UseFilthyTiers = true

			var expectedPicks = make(map[string][]*civ.Civ)
			for k, v := range testcase.expected {
				for _, ck := range v {
					expectedPicks[k] = append(expectedPicks[k], cs.CivMap[ck])
				}
			}

			cs.makePicksWithTier()
			assert.EqualValues(t, expectedPicks, cs.Picks)

			for _, v := range cs.Picks {
				assert.True(t, v[0].FilthyTier == 1 || v[0].FilthyTier == 2)
			}
		})
	}
}

func TestMakePicksWithoutTier(t *testing.T) {
	testdata := NewTestData()
	players := testdata.Players
	clk := clock.NewMock()

	testcases := []struct {
		description string
		civs        []constants.CivKey
		players     []*discordgo.User
		numPicks    int
		expected    map[string][]constants.CivKey
	}{
		{
			description: "nil civ list returns nil",
			civs:        nil,
			players:     players[:1],
			numPicks:    1,
			expected:    nil,
		},
		{
			description: "one pick, one player, one civ returns that civ",
			civs:        []constants.CivKey{constants.ARABIA},
			players:     players[:1],
			numPicks:    1,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.ARABIA},
			},
		},
		{
			description: "two picks, one player, two civs returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA},
			players:     players[:1],
			numPicks:    2,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.AMERICA, constants.ARABIA},
			},
		},
		{
			description: "one pick, two players, four civs returns expected civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     players[:2],
			numPicks:    1,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.ARABIA},
				players[1].ID: {constants.AMERICA},
			},
		},
		{
			description: "two picks, two players, four civs returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     players[:2],
			numPicks:    2,
			expected: map[string][]constants.CivKey{
				players[0].ID: {constants.AMERICA, constants.ARABIA},
				players[1].ID: {constants.ASSYRIA, constants.AZTECS},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			// TODO: all of this setup can probably be pulled into a helper.
			cs := NewCivSession()
			cs.Clock = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, cs.CivMap[k])
			}
			civMap := civ.GenCivMap(civsToTest)
			cs.Civs = civsToTest
			cs.CivMap = civMap

			var playerMap = make(map[string]*discordgo.User)
			for _, u := range testcase.players {
				playerMap[u.ID] = u
			}
			cs.Players = testcase.players
			cs.PlayerMap = playerMap

			cs.Config.NumPicks = testcase.numPicks

			var expectedPicks = make(map[string][]*civ.Civ)
			for k, v := range testcase.expected {
				for _, ck := range v {
					expectedPicks[k] = append(expectedPicks[k], cs.CivMap[ck])
				}
			}

			cs.makePicksWithoutTier()
			assert.EqualValues(t, expectedPicks, cs.Picks)
		})
	}
}
