package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Message struct {
	*discordgo.Message
}

type MessageCreate struct {
	Message
}
