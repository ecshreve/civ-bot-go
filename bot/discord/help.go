package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/constants"
)

// helpCommandHandler sends a help message. If a specific command is provided it
// provides that specific help message, otherwise it provides the default summary
// help message for all commands.
func helpCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	topic := ""
	if len(args) > 1 {
		topic = args[1]
	}

	// TODO: this is dumb, should make it better.
	switch topic {
	case "new":
		title = "🆕  new"
		description = "starts a new civ-bot session in the current channel \n whoever wants to play reacts with  ✋\n someone adds a  ✅ react when ready to continue"
	case "oops":
		title = "🤷‍♀️  oops"
		description = "abandon current session and start over"
	case "info":
		title = "ℹ️  info"
		description = "output the current state of the session"
	case "ban":
		title = "🍌  ban"
		description = "ban a civ so it can't be part of a player's picks"
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
				Name:  "🤷‍♀️  oops",
				Value: "`/civ help oops`: abandon current session and start over",
			},
			{
				Name:  "ℹ️  info",
				Value: "`/civ help info`: output the current state of the session",
			},
			{
				Name:  "🍌  ban",
				Value: "`/civ help ban`: ban a civ so it can't be part of a player's picks",
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
		Color:       constants.ColorBLUE,
	})
}
