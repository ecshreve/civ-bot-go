package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/constants"
)

// banInstructions sends a Message with instructions for how to ban a Civ.
func banInstructions(s *discordgo.Session, channelID string) {
	cs := Session
	s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "ℹ️ okay, here's our players",
		Description: "- everyone gets to ban a civ now, enter `/civ ban <civ name>` to choose\n- if you change your mind just enter `/civ ban <new civ name>` to update your choice\n\nnote: you can enter a ban by either the civ or leader name",
		Color:       constants.ColorGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Players",
				Value: formatUsers(cs.Players),
			},
		},
	})
}

// banCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func banCiv(civToBan string, uid string) *Civ {
	cs := Session
	c := cs.getCivByString(civToBan)
	if c == nil || c.Banned == true {
		return nil
	}

	// If this player had previously banned a Civ then unban the previous Civ.
	if _, ok := cs.Bans[uid]; ok {
		cs.Bans[uid].Banned = false
	}

	c.Banned = true
	cs.Bans[uid] = c

	return c
}
