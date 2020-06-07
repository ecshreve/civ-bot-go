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
