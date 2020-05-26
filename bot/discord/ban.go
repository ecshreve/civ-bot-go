package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/constants"
)

// banInstructions sends a Message with instructions for how to ban a Civ.
func (cs *CivSession) banInstructions(s *discordgo.Session, channelID string) {
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
func (cs *CivSession) banCiv(civToBan string, uid string) *Civ {
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

func (cs *CivSession) banCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("ban missing", "🤔  "+formatUser(m.Author)+" you have to actually ban someone"))
		return
	}

	c := cs.banCiv(args[1], m.Author.ID)
	if c == nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("invalid ban", "🤔  "+formatUser(m.Author)+" can you pick a valid civ to ban?"))
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "🍌 current bans",
		Color: constants.ColorRED,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "bans",
				Value: cs.formatBans(),
			},
		},
	})

	// If all players have entered a Ban then pick Civs for all players.
	if len(cs.Bans) == len(cs.Players) {
		cs.pick(s, m)
	}
}
