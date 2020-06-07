package bot

import (
	"github.com/bwmarrin/discordgo"

	"github.com/ecshreve/civ-bot-go/internal/constants"
)

// CommandID is a wrapper around the string representation of the command.
type CommandID string

// CommandInfo defines the properties of a Command.
type CommandInfo struct {
	Name        string
	Emoji       string
	Description string
	Usage       string
}

// Command is any command sent to the Bot.
type Command interface {
	Info() *CommandInfo
	Process(*discordgo.Message) *discordgo.MessageEmbed
}

// Command interface implementation for the "help" command.
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
