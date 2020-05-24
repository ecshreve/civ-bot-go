package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func infoCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, cs *CivSession) {
	title := "ℹ️ current civ session info"
	players := ""

	if len(cs.Players) == 0 {
		players = "no players yet"
	} else {
		for _, p := range cs.Players {
			players = players + formatUser(p) + "\n"
		}
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: title,
		Color: cGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Players",
				Value: players,
			},
		},
	})

	if err != nil {
		fmt.Printf("error generating info: %+v", err)
		return
	}
}
