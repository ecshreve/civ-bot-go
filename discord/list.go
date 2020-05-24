package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (cs *CivSession) listCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "☁︎  list all possible civs",
		Color: cGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "all civs",
				Value: formatCivs(cs.Civs),
			},
		},
	})

	if err != nil {
		fmt.Printf("error listing civs: %+v", err)
		return
	}
}
