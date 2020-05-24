package discord

import "github.com/bwmarrin/discordgo"

func helpCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, topic string) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	switch topic {
	case "new":
		title = "🆕  new - civ-bot help"
		description = "starts a new civ-bot session"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "general new",
				Value: "`/civ new`: starts a new civ-bot session in the current channel",
			},
		}

	default:
		title = "ℹ️  topics - civ-bot"
		description = "pick a topic to get help"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "🆕  new",
				Value: "`/civ help new`:  instructions on starting a new civ-bot session",
			},
		}
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       cBLUE,
	})
}
