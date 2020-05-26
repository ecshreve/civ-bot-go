package discord

import (
	"github.com/bwmarrin/discordgo"
)

// newReactionHandler handles all new related reactions.
func newReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "✋" {
		Session.Players[user.ID] = user
	}
	if r.Emoji.Name == "✅" {
		banInstructions(s, m.ChannelID)
	}
}
