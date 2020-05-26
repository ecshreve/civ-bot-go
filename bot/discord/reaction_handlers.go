package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/civsession"
	"github.com/ecshreve/civ-bot-go/bot/constants"
	"github.com/ecshreve/civ-bot-go/bot/util"
)

// newReactionHandler handles all new related reactions.
func newReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	cs := civsession.CS
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

func pickReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "♻️" {
		civsession.CS.RePickVotes++
	}
}
