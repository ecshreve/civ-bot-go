package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

type helpCommand struct{}

func (c *helpCommand) Info() *CommandInfo {
	return &CommandInfo{
		Name:        "Help",
		Emoji:       "ℹ️",
		Description: "view list of commands",
		Usage:       "/civ help",
	}
}

func (c *helpCommand) Process(m *discordgo.Message) *discordgo.MessageEmbed {
	info := c.Info()
	title := "available bot commands"
	description := "---"
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  info.Emoji + " " + info.Name,
			Value: info.Usage + ": " + info.Description,
		},
	}

	return &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       constants.ColorBLUE,
	}
}
