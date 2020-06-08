package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/ecshreve/civ-bot-go/internal/util"
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
	Process(*Bot, *discordgo.Message) (*discordgo.Message, error)
}

// AllCommands is a slice containing all the valid Commands for the Bot. It's mainly
// for convenience so we can maintain parity between the actual Bot implementation
// and the MockBot we use for testing.
var AllCommands = []Command{
	&helpCommand{},
	&configCommand{},
	&newCommand{},
	&oopsCommand{},
	&banCommand{},
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

func (c *helpCommand) Process(b *Bot, m *discordgo.Message) (*discordgo.Message, error) {
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

	helpMessage, err := b.DS.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return nil, oops.Wrapf(err, "error sending embed %+v", embed)
	}

	return helpMessage, nil
}

// Command interface implementation for the "config" command.
type configCommand struct{}

func (c *configCommand) Info() *CommandInfo {
	return &CommandInfo{
		Name:        "Config",
		Emoji:       "⚙️",
		Description: "enter the session configuration menu",
		Usage:       "`/civ config`",
	}
}

func (c *configCommand) Process(b *Bot, m *discordgo.Message) (*discordgo.Message, error) {
	title := "⚙️ configuration"
	description := "- select 🛠 to change config\n- select ✅ to accept config"
	fields := b.CivConfig.GetEmbedFields()

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       constants.ColorDARKGREY,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "config",
		},
	}

	configMessage, err := b.DS.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return nil, oops.Wrapf(err, "error sending embed: %+v", embed)
	}

	err = b.DS.MessageReactionAdd(m.ChannelID, configMessage.ID, "🛠")
	if err != nil {
		return configMessage, oops.Wrapf(err, "unable to add reaction %s to embed: %+v", "✋", embed)
	}

	err = b.DS.MessageReactionAdd(m.ChannelID, configMessage.ID, "✅")
	if err != nil {
		return configMessage, oops.Wrapf(err, "unable to add reaction %s to embed: %+v", "✅", embed)
	}

	return configMessage, nil
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

func (c *newCommand) Process(b *Bot, m *discordgo.Message) (*discordgo.Message, error) {
	embedTitle := "🆕 starting a new civ bot session"
	newMessage, err := newSessionHelper(b, m.ChannelID, embedTitle)
	if err != nil {
		return newMessage, oops.Wrapf(err, "unable to process newCommand")
	}

	b.Reset(false)
	return newMessage, nil
}

// Command interface implementation for the "oops" command.
type oopsCommand struct{}

func (c *oopsCommand) Info() *CommandInfo {
	return &CommandInfo{
		Name:        "Oops",
		Emoji:       "🤷‍♀️",
		Description: "start a new civ-bot session maintaining the current config",
		Usage:       "`/civ oops`",
	}
}

func (c *oopsCommand) Process(b *Bot, m *discordgo.Message) (*discordgo.Message, error) {
	embedTitle := "🤷‍♀️restarting the civ bot session with the same config"
	oopsMessage, err := newSessionHelper(b, m.ChannelID, embedTitle)
	if err != nil {
		return oopsMessage, oops.Wrapf(err, "unable to process oopsCommand")
	}

	b.Reset(true)
	return oopsMessage, nil
}

// Command interface implementation for the "ban" command.
type banCommand struct{}

func (c *banCommand) Info() *CommandInfo {
	return &CommandInfo{
		Name:        "Ban",
		Emoji:       "🍌",
		Description: "ban the Civ that most closely matches the string argument",
		Usage:       "`/civ ban <string>`",
	}
}

func (c *banCommand) Process(b *Bot, m *discordgo.Message) (*discordgo.Message, error) {
	// If a player not in the Bot session tries to ban then return an error message.
	player, ok := b.CivState.PlayerMap[PlayerID(m.Author.ID)]
	if !ok {
		return b.DS.ChannelMessageSend(m.ChannelID, util.ErrorMessage("uhhh... ", "yo <@"+m.Author.ID+"> you aren't in this game, enter `/civ oops` to restart the session"))
	}

	// If no Civ was provided then return an error message.
	args := strings.Split(m.Content, " ")[1:]
	if len(args) == 1 {
		return b.DS.ChannelMessageSend(m.ChannelID, util.ErrorMessage("ban missing", "🤔  "+FormatPlayer(player)+" you have to actually ban someone"))
	}

	_, err := player.BanCiv(b, args[1])
	if err != nil {
		return b.DS.ChannelMessageSend(m.ChannelID, util.ErrorMessage("invalid ban", "🤔  "+FormatPlayer(player)+" can you pick a valid civ to ban?"))
	}

	var embedFields []*discordgo.MessageEmbedField
	for k, v := range b.CivState.Bans {
		f := &discordgo.MessageEmbedField{
			Name:  b.CivState.PlayerMap[k].Username,
			Value: civ.FormatCivs(v),
		}
		embedFields = append(embedFields, f)
	}
	embed := &discordgo.MessageEmbed{
		Title:  "🍌 current bans",
		Color:  constants.ColorRED,
		Fields: embedFields,
	}

	banMessage, err := b.DS.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return nil, oops.Wrapf(err, "error sending embed: %+v", embed)
	}

	return banMessage, nil
}

// newSessionHelper is called when we Process either a newCommand or an oopsCommand.
// Pulling this out into a function lets us handle resetting the CivState and CivConfig
// differently depending on which command is processed.
func newSessionHelper(b *Bot, channelID, embedTitle string) (*discordgo.Message, error) {
	description := "- whoever wants to play react with  ✋\n- someone add a  ✅ react when ready to continue \n\n- enter `/civ config` to view or update the configuration \n- enter `/civ oops` at any point to completely start over\n- enter `/civ help` to see a list of available commands"

	embed := &discordgo.MessageEmbed{
		Title:       embedTitle,
		Description: description,
		Color:       constants.ColorDARKPURPLE,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "new",
		},
	}

	newMessage, err := b.DS.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return nil, oops.Wrapf(err, "error sending embed: %+v", embed)
	}

	err = b.DS.MessageReactionAdd(channelID, newMessage.ID, "✋")
	if err != nil {
		return newMessage, oops.Wrapf(err, "unable to add reaction %s to embed: %+v", "✋", embed)
	}

	err = b.DS.MessageReactionAdd(channelID, newMessage.ID, "✅")
	if err != nil {
		return newMessage, oops.Wrapf(err, "unable to add reaction %s to embed: %+v", "✅", embed)
	}

	return newMessage, nil
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
