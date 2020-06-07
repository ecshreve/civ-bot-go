package bot

import (
	"github.com/bwmarrin/discordgo"
)

type MessageInterface interface{}

type MessCreate struct {
	*discordgo.MessageCreate
}

type EmbedField struct {
	*discordgo.MessageEmbedField
}

func NewEmbedField(name, value string) *EmbedField {
	dgEmbedField := &discordgo.MessageEmbedField{
		Name:  name,
		Value: value,
	}

	ef := &EmbedField{dgEmbedField}
	return ef
}
