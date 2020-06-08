package bot

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/samsarahq/go/oops"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
)

func TestMakePick(t *testing.T) {
	clk := clock.NewMock()

	testcases := []struct {
		description       string
		civs              []constants.CivKey
		initialBannedCivs []constants.CivKey
		initialPickedCivs []constants.CivKey
		expectError       bool
		expected          constants.CivKey
	}{
		{
			description: "empty input slice returns error",
			civs:        []constants.CivKey{},
			expectError: true,
		},
		{
			description:       "pick from slice of length 1 where that civ is already picked returns error",
			civs:              []constants.CivKey{constants.AMERICA},
			initialPickedCivs: []constants.CivKey{constants.AMERICA},
			expectError:       true,
		},
		{
			description:       "pick from slice of length 1 where that civ is banned returns error",
			civs:              []constants.CivKey{constants.AMERICA},
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			expectError:       true,
		},
		{
			description:       "pick from slice of length 2 where one civ is banned, one is picked returns error",
			civs:              []constants.CivKey{constants.AMERICA, constants.ARABIA},
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			initialPickedCivs: []constants.CivKey{constants.ARABIA},
			expectError:       true,
		},
		{
			description: "pick from slice of length 1 returns that item",
			civs:        []constants.CivKey{constants.AMERICA},
			expected:    constants.AMERICA,
		},
		{
			description:       "pick from slice of length 2 where one civ is banned returns the other civ",
			civs:              []constants.CivKey{constants.AMERICA, constants.ARABIA},
			initialBannedCivs: []constants.CivKey{constants.AMERICA},
			expected:          constants.ARABIA,
		},
		// This testcase is deterministic because we provide a MockClock to our
		// CivSession, and we seed our random number generator in MakePick with
		// the current time.
		{
			description: "pick from slice of all civs returns indonesia",
			civs:        constants.CivKeys,
			expected:    constants.INDONESIA,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			b, _ := MockBot(t)
			b.CivState.Clk = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, b.CivState.CivMap[k])
			}

			for _, k := range testcase.initialBannedCivs {
				b.CivState.CivMap[k].Banned = true
			}

			for _, k := range testcase.initialPickedCivs {
				b.CivState.CivMap[k].Picked = true
			}

			actual, err := b.makePick(civsToTest)
			if !testcase.expectError {
				assert.NoError(t, err)
				assert.Equal(t, b.CivState.CivMap[testcase.expected], actual)
				assert.Same(t, b.CivState.CivMap[testcase.expected], actual)
			} else {
				assert.Nil(t, actual)
				assert.Error(t, err)
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
		expectError       bool
		expected          []constants.CivKey
	}{
		{
			description: "empty input slice returns error",
			civs:        []constants.CivKey{},
			numPicks:    1,
			expectError: true,
		},
		{
			description: "input slice less than numPicks returns error",
			civs:        []constants.CivKey{constants.AMERICA},
			numPicks:    2,
			expectError: true,
		},
		{
			description:       "picking 2 civs from list of 2 with one already picked returns error",
			civs:              []constants.CivKey{constants.AMERICA, constants.ARABIA},
			numPicks:          2,
			initialPickedCivs: []constants.CivKey{constants.AMERICA},
			expectError:       true,
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
			b, _ := MockBot(t)
			b.CivState.Clk = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, b.CivState.CivMap[k])
			}

			for _, k := range testcase.initialPickedCivs {
				b.CivState.CivMap[k].Picked = true
			}

			var expectedCivs []*civ.Civ
			for _, k := range testcase.expected {
				expectedCivs = append(expectedCivs, b.CivState.CivMap[k])
			}

			actual, err := b.makePicks(civsToTest, testcase.numPicks)
			if testcase.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.EqualValues(t, expectedCivs, actual)
		})
	}
}

func TestMakePicksWithTier(t *testing.T) {
	clk := clock.NewMock()

	testUserIDs := []string{"testPlayer1", "testPlayer2", "testPlayer2"}
	var testPlayers []*Player
	for _, id := range testUserIDs {
		testUser := &discordgo.User{ID: id}
		testPlayer := NewPlayer(testUser)
		testPlayers = append(testPlayers, testPlayer)
	}

	testcases := []struct {
		description string
		civs        []constants.CivKey
		players     []*Player
		numPicks    int
		expectError bool
		expected    map[PlayerID][]constants.CivKey
	}{
		{
			description: "nil civ list returns error",
			civs:        nil,
			players:     testPlayers[:1],
			numPicks:    1,
			expectError: true,
		},
		{
			description: "one pick, one player, one low tier civ returns error",
			civs:        []constants.CivKey{constants.AMERICA},
			players:     testPlayers[:1],
			numPicks:    1,
			expectError: true,
		},
		{
			description: "one pick, one player, one top tier civ returns that civ",
			civs:        []constants.CivKey{constants.ARABIA},
			players:     testPlayers[:1],
			numPicks:    1,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.ARABIA},
			},
		},
		{
			description: "two picks, one player, one top tier civ, zero low tier civ returns error",
			civs:        []constants.CivKey{constants.ARABIA},
			players:     testPlayers[:1],
			numPicks:    2,
			expectError: true,
		},
		{
			description: "two picks, one player, one top tier civ, one low tier civ returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA},
			players:     testPlayers[:1],
			numPicks:    2,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.ARABIA, constants.AMERICA},
			},
		},
		{
			description: "one pick, two players, two top tier civ, two low tier civ returns the top tier civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     testPlayers[:2],
			numPicks:    1,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.AZTECS},
				testPlayers[1].PlayerID: {constants.ARABIA},
			},
		},
		{
			description: "two picks, two players, two top tier civ, two low tier civ returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     testPlayers[:2],
			numPicks:    2,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.AZTECS, constants.AMERICA},
				testPlayers[1].PlayerID: {constants.ARABIA, constants.ASSYRIA},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			// TODO: all of this setup can probably be pulled into a helper.
			b, _ := MockBot(t)
			cs := b.CivState
			cs.Clk = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, cs.CivMap[k])
			}
			civMap := civ.GenCivMap(civsToTest)
			cs.Civs = civsToTest
			cs.CivMap = civMap

			var playerMap = make(map[PlayerID]*Player)
			for _, p := range testcase.players {
				playerMap[p.PlayerID] = p
			}
			cs.Players = testcase.players
			cs.PlayerMap = playerMap

			b.CivConfig.Picks = testcase.numPicks
			b.CivConfig.UseTiers = true

			var expectedPicks = make(map[PlayerID][]*civ.Civ)
			for k, v := range testcase.expected {
				for _, ck := range v {
					expectedPicks[k] = append(expectedPicks[k], cs.CivMap[ck])
				}
			}

			err := b.makePicksWithTier()
			if testcase.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.EqualValues(t, expectedPicks, cs.Picks)
			for _, v := range cs.Picks {
				assert.True(t, v[0].FilthyTier == 1 || v[0].FilthyTier == 2)
			}
		})
	}
}

func TestMakePicksWithoutTier(t *testing.T) {
	clk := clock.NewMock()

	testUserIDs := []string{"testPlayer1", "testPlayer2", "testPlayer2"}
	var testPlayers []*Player
	for _, id := range testUserIDs {
		testUser := &discordgo.User{ID: id}
		testPlayer := NewPlayer(testUser)
		testPlayers = append(testPlayers, testPlayer)
	}

	testcases := []struct {
		description string
		civs        []constants.CivKey
		players     []*Player
		numPicks    int
		expectError bool
		expected    map[PlayerID][]constants.CivKey
	}{
		{
			description: "nil civ list returns error",
			civs:        nil,
			players:     testPlayers[:1],
			numPicks:    1,
			expectError: true,
		},
		{
			description: "one pick, one player, one civ returns that civ",
			civs:        []constants.CivKey{constants.ARABIA},
			players:     testPlayers[:1],
			numPicks:    1,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.ARABIA},
			},
		},
		{
			description: "two picks, one player, two civs returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA},
			players:     testPlayers[:1],
			numPicks:    2,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.AMERICA, constants.ARABIA},
			},
		},
		{
			description: "one pick, two players, four civs returns expected civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     testPlayers[:2],
			numPicks:    1,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.ARABIA},
				testPlayers[1].PlayerID: {constants.AMERICA},
			},
		},
		{
			description: "two picks, two players, four civs returns those civs",
			civs:        []constants.CivKey{constants.AMERICA, constants.ARABIA, constants.ASSYRIA, constants.AZTECS},
			players:     testPlayers[:2],
			numPicks:    2,
			expected: map[PlayerID][]constants.CivKey{
				testPlayers[0].PlayerID: {constants.AMERICA, constants.ARABIA},
				testPlayers[1].PlayerID: {constants.ASSYRIA, constants.AZTECS},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			// TODO: all of this setup can probably be pulled into a helper.
			b, _ := MockBot(t)
			cs := b.CivState
			cs.Clk = clk

			var civsToTest []*civ.Civ
			for _, k := range testcase.civs {
				civsToTest = append(civsToTest, cs.CivMap[k])
			}
			civMap := civ.GenCivMap(civsToTest)
			cs.Civs = civsToTest
			cs.CivMap = civMap

			var playerMap = make(map[PlayerID]*Player)
			for _, p := range testcase.players {
				playerMap[p.PlayerID] = p
			}
			cs.Players = testcase.players
			cs.PlayerMap = playerMap

			b.CivConfig.Picks = testcase.numPicks

			var expectedPicks = make(map[PlayerID][]*civ.Civ)
			for k, v := range testcase.expected {
				for _, ck := range v {
					expectedPicks[k] = append(expectedPicks[k], cs.CivMap[ck])
				}
			}

			err := b.makePicksWithoutTier()
			if testcase.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.EqualValues(t, expectedPicks, cs.Picks)
		})
	}
}

func TestPick(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	clk := clock.NewMock()
	b, mock := MockBot(t)
	b.CivState.Clk = clk
	testChannelID := "testChannel"

	testUserIDs := []string{"testPlayer1", "testPlayer2", "testPlayer3"}
	var testPlayers []*Player
	for _, id := range testUserIDs {
		testUser := &discordgo.User{ID: id}
		testPlayer := NewPlayer(testUser)
		testPlayers = append(testPlayers, testPlayer)
	}
	b.CivState.Players = testPlayers
	b.CivState.PlayerMap = GetPlayerIDToPlayerMap(testPlayers)
	b.CivState.RePicksRemaining = 1

	// Kick off a ticker to increment every 1 mock second.
	go func() {
		ticker := clk.Ticker(1 * time.Second)
		for {
			<-ticker.C
			err := b.Pick(testChannelID)
			assert.NoError(t, err)
			return
		}
	}()
	runtime.Gosched()

	// Expect the bot to add the initial message and reaction.
	mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, MockAny{})
	mock.Expect(b.DS.MessageReactionAdd, testChannelID, MockAny{}, "♻️")

	// Expect the message to be edited every seconds for 60 seconds.
	for i := 0; i < 60; i++ {
		mock.Expect(b.DS.ChannelMessageEditEmbed, testChannelID, MockAny{}, MockAny{})
	}

	// Expect the bot to send the final message.
	sessionOverEmbed := &discordgo.MessageEmbed{
		Title: "great, have fun! see y'all next time 👋",
		Color: constants.ColorORANGE,
	}
	mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, sessionOverEmbed)
	clk.Add(61 * time.Second)
	snap.Snapshot("picks", b.CivState.Picks)

	// Reset the CivState and set RePicksRemaining to 0.
	b.CivState = NewCivState()
	b.CivState.Clk = clk
	b.CivState.Players = testPlayers
	b.CivState.PlayerMap = GetPlayerIDToPlayerMap(testPlayers)

	// Do the same thing but with RePicks set to 0.
	b.CivState.RePicksRemaining = 0

	// Expect the bot to add the initial.
	mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, MockAny{})

	// Expect the bot to send the final "no repicks remaining" message.
	sessionOverEmbed = &discordgo.MessageEmbed{
		Title: "no more re-picks, those are your choices, deal with it",
		Color: constants.ColorORANGE,
	}
	mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, sessionOverEmbed)
	b.Pick(testChannelID)
	snap.Snapshot("picks2", b.CivState.Picks)

	// Reset the CivState.
	b.CivState = NewCivState()
	b.CivState.Clk = clk
	b.CivState.Players = testPlayers
	b.CivState.PlayerMap = GetPlayerIDToPlayerMap(testPlayers)

	// CivConfig.UseTiers false and empty CivState.Civs should result in error.
	b.CivState.Civs = nil
	b.CivState.CivMap = nil
	err := b.Pick(testChannelID)
	assert.Error(t, err)
	snap.Snapshot("UseTiers false error", fmt.Sprintf("%v", oops.Cause(err).Error()))

	// CivConfig.UseTiers true and empty CivState.Civs should result in error.
	b.CivConfig.UseTiers = true
	err = b.Pick(testChannelID)
	assert.Error(t, err)
	snap.Snapshot("UseTiers true error", fmt.Sprintf("%v", oops.Cause(err).Error()))

	Check(mock.Check(), true, t)
}

func TestCountdown(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	clk := clock.NewMock()
	b, mock := MockBot(t)
	b.CivState.Clk = clk
	testStart := clk.Now()

	embedFields := []*discordgo.MessageEmbedField{
		{
			Name:  "name1",
			Value: "value1",
		},
		{
			Name:  "name2",
			Value: "value2",
		},
	}

	pickEmbed := &discordgo.MessageEmbed{
		Title:       "picks",
		Description: "pick description for test",
		Color:       constants.ColorDARKBLUE,
		Fields:      embedFields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "pick",
		},
	}

	pickMessage := &discordgo.Message{
		ID:        "messageID",
		ChannelID: "testChannelID",
	}

	err := b.countdown(pickMessage, testStart, 10)
	assert.Error(t, err)
	snap.Snapshot("0 embeds", fmt.Sprintf(oops.Cause(err).Error()))

	pickMessage.Embeds = []*discordgo.MessageEmbed{pickEmbed, pickEmbed}
	err = b.countdown(pickMessage, testStart, 10)
	assert.Error(t, err)
	snap.Snapshot("2 embeds", fmt.Sprintf(oops.Cause(err).Error()))

	pickMessage.Embeds = []*discordgo.MessageEmbed{pickEmbed}
	go func() {
		for {
			err = b.countdown(pickMessage, testStart, 10)
			assert.NoError(t, err)
		}
	}()
	runtime.Gosched()

	// Countdown from 10 to 5 and snapshot the Message and CivState.
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	clk.Add(5 * time.Second)
	snap.Snapshot("pickMessage before timer runs out", pickMessage)
	snap.Snapshot("PickState before timer runs out", b.CivState.PickState)

	// Countdown from 5 to 0 and snapshot the Message.
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageSendEmbed, pickMessage.ChannelID, MockAny{})
	clk.Add(5 * time.Second)
	snap.Snapshot("pickMessage when timer runs out", pickMessage)
	snap.Snapshot("PickState when timer runs out", b.CivState.PickState)

	// Reset the clock.
	clk = clock.NewMock()
	b.CivState.Clk = clk
	go func() {
		for {
			err = b.countdown(pickMessage, testStart, 10)
			assert.NoError(t, err)
		}
	}()
	runtime.Gosched()

	// Countdown from 10 to 5.
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	clk.Add(5 * time.Second)

	// Set DoRepick to true and countdown from 5 to 4.
	b.CivState.DoRepick = true
	mock.Expect(b.DS.ChannelMessageEditEmbed, pickMessage.ChannelID, pickMessage.ID, MockAny{})
	mock.Expect(b.DS.ChannelMessageSendEmbed, pickMessage.ChannelID, MockAny{})
	clk.Add(1 * time.Second)
	snap.Snapshot("picks remaining decremented to 2", b.CivState.PickState)

	Check(mock.Check(), true, t)
}
