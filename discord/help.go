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
				Name:  "basic operation",
				Value: "`/civ new`: starts a new civ-bot session in the current channel \n whoever wants to play reacts with  ✋\n someone adds a  ✅ react when ready to continue",
			},
		}
	case "list":
		title = "☁︎  list"
		description = "lists all civs and leaders"
	default:
		title = "ℹ️  topics - civ-bot"
		description = "pick a topic to get help"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "🆕  new",
				Value: "`/civ help new`: instructions on starting a new civ-bot session",
			},
			{
				Name:  "☁︎  list",
				Value: "`/civ help list`: lists all possible civs",
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
