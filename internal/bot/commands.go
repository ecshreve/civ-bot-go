package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"

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
	Process(*Bot, *discordgo.Message) error
}

// Command interface implementation for the "help" command.
type helpCommand struct{}

func (c *helpCommand) Info() *CommandInfo {
	return &CommandInfo{
		Name:        "Help",
		Emoji:       "ℹ️",
		Description: "view list of commands",
		Usage:       "`/civ help`",
	}
}

func (c *helpCommand) Process(b *Bot, m *discordgo.Message) error {
	title := "available bot commands"
	description := "---"

	var fields []*discordgo.MessageEmbedField
	for _, c := range b.Commands {
		fields = append(fields, getHelpEmbedField(c))
	}

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       constants.ColorBLUE,
	}

	_, err := b.DS.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return oops.Wrapf(err, "error sending embed %+v", embed)
	}

	return nil
}

// Command interface implementation for the "new" command.
type newCommand struct{}

func (c *newCommand) Info() *CommandInfo {
	return &CommandInfo{
		Name:        "New",
		Emoji:       "🆕",
		Description: "start a new civ-bot session",
		Usage:       "`/civ new`",
	}
}

func (c *newCommand) Process(b *Bot, m *discordgo.Message) error {
	b.CivState.Reset(b.CivConfig)

	title := "🆕 starting a new civ picker session"
	description := "- whoever wants to play react with  ✋\n- someone add a  ✅ react when ready to continue \n\n- enter `/civ config` to view or update the configuration \n- enter `/civ oops` at any point to completely start over\n- enter `/civ help` to see a list of available commands"

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       constants.ColorDARKPURPLE,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "new",
		},
	}

	newEmbed, err := b.DS.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return oops.Wrapf(err, "error sending embed: %+v", embed)
	}

	err = b.DS.MessageReactionAdd(m.ChannelID, newEmbed.ID, "✋")
	if err != nil {
		return oops.Wrapf(err, "unable to add reaction %s to embed: %+v", "✋", embed)
	}

	err = b.DS.MessageReactionAdd(m.ChannelID, newEmbed.ID, "✅")
	if err != nil {
		return oops.Wrapf(err, "unable to add reaction %s to embed: %+v", "✅", embed)
	}

	return nil
}

// getHelpEmbedField returns a MessageEmbedField for the given Command that's used
// when processing a helpCommand.
func getHelpEmbedField(c Command) *discordgo.MessageEmbedField {
	info := c.Info()

	return &discordgo.MessageEmbedField{
		Name:  info.Emoji + " " + info.Name,
		Value: info.Usage + ": " + info.Description,
	}
}

// getCommandIDToCommandMap returns a map of [CommandID]Command for the given slice of Commands.
func getCommandIDToCommandMap(commands []Command) map[CommandID]Command {
	commandMap := make(map[CommandID]Command)
	for _, c := range commands {
		commandMap[CommandID(strings.ToLower(c.Info().Name))] = c
	}

	return commandMap
}
