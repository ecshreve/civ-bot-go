package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/samsarahq/go/snapshotter"
)

func TestNewReaction(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, mock := MockBot(t)
	new := b.ReactionMap["new"]

	// Emoji should add a Player to the CivState.
	testReaction := &discordgo.MessageReaction{
		UserID:    "testUser",
		ChannelID: "testChannel",
		Emoji: discordgo.Emoji{
			Name: "✋",
		},
	}

	mock.Expect(b.DS.User, testReaction.UserID)
	newReactionMessage, err := new.Process(b, testReaction)
	assert.NoError(t, err)
	assert.Nil(t, newReactionMessage)
	assert.Equal(t, 1, len(b.CivState.Players))
	assert.Equal(t, 1, len(b.CivState.PlayerMap))

	// Emoji should send a Message with instructions on picking Bans.
	testReaction2 := &discordgo.MessageReaction{
		UserID:    "testUser",
		ChannelID: "testChannel",
		Emoji: discordgo.Emoji{
			Name: "✅",
		},
	}

	mock.Expect(b.DS.ChannelMessageSendEmbed, testReaction2.ChannelID, MockAny{})
	newReactionMessage, err = new.Process(b, testReaction2)
	assert.NoError(t, err)
	snap.Snapshot("message embed after confirming", newReactionMessage.Embeds)
}
