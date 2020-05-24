package discord

import (
	"github.com/bwmarrin/discordgo"
)

// banInstructions sends a Message with instructions for how to ban a Civ.
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

// banCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func (cs *CivSession) banCiv(civToBan string, u *discordgo.User) *Civ {
	c := cs.getCivByString(civToBan)
	if c == nil {
		return nil
	}

	c.Banned = true
	cs.Bans[u] = c
	return c
}

func (cs *CivSession) banCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("ban missing", "🤔  "+formatUser(m.Author)+" you have to actually ban someone"))
		return
	}

	c := cs.banCiv(args[1], m.Author)
	if c == nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("invalid ban", "🤔  "+formatUser(m.Author)+" can you pick a valid civ to ban?"))
		return
	}

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

	// If all players have entered a Ban then pick Civs for all players.
	if len(cs.Bans) == len(cs.Players) {
		cs.pick(s, m)
	}
}
