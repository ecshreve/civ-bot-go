package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (cs *CivSession) banInstructions(s *discordgo.Session, channelID string) {
	title := "ℹ️ okay, here's our players \neveryone gets to ban a civ now, enter `/civ ban <civ name>` to choose"
	s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title: title,
		Color: cGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Players",
				Value: formatUsers(cs.Players),
			},
		},
	})
}

func (cs *CivSession) banCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("ban missing", "🤔  "+formatUser(m.Author)+" you have to actually ban someone"))
		return
	}

	c := cs.getCivByString(args[1])
	if c == nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("invalid ban", "🤔  "+formatUser(m.Author)+" can you pick a valid civ to ban?"))
		return
	}

	c.Banned = true
	cs.Bans[m.Author] = c

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "🍌 current bans",
		Color: cRED,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "bans",
				Value: formatBans(cs.Bans),
			},
		},
	})

	if len(cs.Bans) == len(cs.Players) {
		cs.pick(s, m)
	}
}
