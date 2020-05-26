package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/constants"
)

func (cs *CivSession) listCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var fields []*discordgo.MessageEmbedField
	for _, c := range cs.Civs {
		f := &discordgo.MessageEmbedField{
			Name:  c.CivBase + " -- " + c.LeaderBase,
			Value: fmt.Sprintf("[zigzag guide >>](%s)\n", c.ZigURL),
		}
		fields = append(fields, f)
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  "☁︎  list all possible civs",
		Color:  constants.ColorGREEN,
		Fields: fields,
	})

	if err != nil {
		fmt.Printf("error listing civs: %+v", err)
		return
	}
}
