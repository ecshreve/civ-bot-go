package bot

import (
	"github.com/bwmarrin/discordgo"
)

type CommandID string

type CommandInfo struct {
	Name        string
	Emoji       string
	Description string
	Usage       string
}
type Command interface {
	Info() *CommandInfo
	Process(*discordgo.Message) *discordgo.MessageEmbed
}
