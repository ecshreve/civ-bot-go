package civsession

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/ecshreve/civ-bot-go/internal/util"
)

// newReactionHandler handles all new related reactions.
func (cs *CivSession) newReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "✋" {
		cs.Players[user.ID] = user
	}
	if r.Emoji.Name == "✅" {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "ℹ️ okay, now that we have our players",
			Description: "- everyone gets to ban a civ, enter `/civ ban <civ name>` to choose\n- if you change your mind just enter `/civ ban <new civ name>` to update your choice\n\nnote: you can enter a ban by either the civ or leader name",
			Color:       constants.ColorGREEN,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Players",
					Value: util.FormatUsers(cs.Players),
				},
			},
		})
	}
}

func (cs *CivSession) configReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	embed := m.Embeds[0]
	if r.Emoji.Name == "✅" {
		embed.Description = "✅ **starting civ picker session with the current config** ✅"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
	}
	if r.Emoji.Name == "🛠" {
		embed.Description = "updating config"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
	}
}

func (cs *CivSession) pickReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "♻️" {
		cs.RePickVotes++
	}
}
