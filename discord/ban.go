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

func (cs *CivSession) getBans() map[*discordgo.User]*Civ {
	ret := make(map[*discordgo.User]*Civ, 0)
	for _, c := range cs.Civs {
		if c.Banned != nil {
			ret[c.Banned] = c
		}
	}
	return ret
}

func banCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, cs *CivSession, inp string) {
	c := cs.getCivByString(inp)
	if c == nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("invalid ban", "🤔  "+formatUser(m.Author)+" can you pick a valid civ to ban?"))
		return
	}
	c.Banned = m.Author

	bans := cs.getBans()
	bansStr := formatBans(bans)

	title := "🍌 current bans"

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: title,
		Color: cRED,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "bans",
				Value: bansStr,
			},
		},
	})

	if len(bans) == len(cs.Players) {
		cs.pick(s, m)
	}
}
