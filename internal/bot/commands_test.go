package bot

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
)

func commandTestHelper(t *testing.T, snap *snapshotter.Snapshotter, c Command, b *Bot, m *discordgo.Message) {
	infoResponse := c.Info()
	snap.Snapshot(fmt.Sprintf("%s -- Info()", infoResponse.Name), infoResponse)

	processResponse, err := c.Process(b, m)
	assert.NoError(t, err)
	snap.Snapshot(fmt.Sprintf("%s -- Process() embeds", infoResponse.Name), processResponse.Embeds)

	snap.Snapshot(fmt.Sprintf("%s -- Bot CivConfig", infoResponse.Name), b.CivConfig)
}

func TestConfigCommand(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, mock := MockBot(t)
	config := b.CommandMap["config"].(*configCommand)

	testMessage := &discordgo.Message{
		ChannelID: "testChannelID",
	}
	mock.Expect(b.DS.ChannelMessageSendEmbed, testMessage.ChannelID, MockAny{})
	mock.Expect(b.DS.MessageReactionAdd, testMessage.ChannelID, MockAny{}, "🛠")
	mock.Expect(b.DS.MessageReactionAdd, testMessage.ChannelID, MockAny{}, "✅")

	commandTestHelper(t, snap, config, b, testMessage)
}

func TestHelpCommand(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, mock := MockBot(t)
	help := b.CommandMap["help"].(*helpCommand)

	testMessage := &discordgo.Message{
		ChannelID: "testChannelID",
	}
	mock.Expect(b.DS.ChannelMessageSendEmbed, testMessage.ChannelID, MockAny{})

	commandTestHelper(t, snap, help, b, testMessage)
}

func TestNewCommand(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, mock := MockBot(t)
	new := b.CommandMap["new"].(*newCommand)

	testMessage := &discordgo.Message{
		ChannelID: "testChannelID",
	}
	mock.Expect(b.DS.ChannelMessageSendEmbed, testMessage.ChannelID, MockAny{})
	mock.Expect(b.DS.MessageReactionAdd, testMessage.ChannelID, MockAny{}, "✋")
	mock.Expect(b.DS.MessageReactionAdd, testMessage.ChannelID, MockAny{}, "✅")

	commandTestHelper(t, snap, new, b, testMessage)
}

func TestOopsCommand(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, mock := MockBot(t)
	oops := b.CommandMap["oops"].(*oopsCommand)

	testMessage := &discordgo.Message{
		ChannelID: "testChannelID",
	}
	mock.Expect(b.DS.ChannelMessageSendEmbed, testMessage.ChannelID, MockAny{})
	mock.Expect(b.DS.MessageReactionAdd, testMessage.ChannelID, MockAny{}, "✋")
	mock.Expect(b.DS.MessageReactionAdd, testMessage.ChannelID, MockAny{}, "✅")

	// Set the config values to something other than the default to verify that the
	// config is maintained.
	b.CivConfig.Bans = 77
	b.CivConfig.Picks = 88
	b.CivConfig.RePicks = 99

	commandTestHelper(t, snap, oops, b, testMessage)
}

func TestBanCommand(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, mock := MockBot(t)
	ban := b.CommandMap["ban"].(*banCommand)
	testChannelID := "testChannelID"

	var testUsers []*discordgo.User
	testUserIDs := []string{"testPlayer1", "testPlayer2"}
	for _, id := range testUserIDs {
		testUser := &discordgo.User{ID: id}
		testUsers = append(testUsers, testUser)

		testPlayer := NewPlayer(testUser)
		b.CivState.Players = append(b.CivState.Players, testPlayer)
		b.CivState.PlayerMap[testPlayer.PlayerID] = testPlayer
	}

	testcases := []struct {
		description string
		content     string
		author      *discordgo.User
		setMocks    func()
	}{
		{
			description: "no civ provided",
			content:     "/civ ban",
			author:      testUsers[0],
			setMocks: func() {
				mock.Expect(b.DS.ChannelMessageSend, testChannelID, MockAny{})
			},
		},
		{
			description: "ban civ",
			content:     "/civ ban america",
			author:      testUsers[0],
			setMocks: func() {
				mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, MockAny{})
			},
		},
		{
			description: "ban an already banned civ",
			content:     "/civ ban america",
			author:      testUsers[0],
			setMocks: func() {
				mock.Expect(b.DS.ChannelMessageSend, testChannelID, MockAny{})
			},
		},
		{
			description: "ban a different civ",
			content:     "/civ ban korea",
			author:      testUsers[0],
			setMocks: func() {
				mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, MockAny{})
			},
		},
		{
			description: "different player ban a different civ",
			content:     "/civ ban america",
			author:      testUsers[1],
			setMocks: func() {
				mock.Expect(b.DS.ChannelMessageSendEmbed, testChannelID, MockAny{})
			},
		},
		{
			description: "player not in session tries to ban",
			content:     "/civ ban brazil",
			author:      &discordgo.User{ID: "testInvalidUser"},
			setMocks: func() {
				mock.Expect(b.DS.ChannelMessageSend, testChannelID, MockAny{})
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			testMessage := &discordgo.Message{
				ChannelID: testChannelID,
				Author:    testcase.author,
				Content:   testcase.content,
			}

			testcase.setMocks()
			banMessage, err := ban.Process(b, testMessage)
			assert.NoError(t, err)
			Check(mock.Check(), true, t)

			snap.Snapshot(fmt.Sprintf("%s - message content", testcase.description), banMessage.Content)
			snap.Snapshot(fmt.Sprintf("%s - message embeds", testcase.description), banMessage.Embeds)
			snap.Snapshot(fmt.Sprintf("%s - civ state bans", testcase.description), b.CivState.Bans)
		})
	}

}
