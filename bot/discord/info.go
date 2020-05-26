package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/constants"
)

func (cs *CivSession) infoCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	title := "ℹ️ current civ session info"
	players := formatUsers(cs.Players)
	if players == "" {
		players = "no players yet"
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: title,
		Color: constants.ColorGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "players",
				Value: players,
			},
			{
				Name:  "bans",
				Value: cs.formatBans(),
			},
		},
	})

	if err != nil {
		fmt.Printf("error generating info: %+v", err)
		return
	}
}
