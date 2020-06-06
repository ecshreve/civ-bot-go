package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"
)

type DataAccessLayer interface {
	Open() error
	Close() error

	ChannelMessage(channelID string, messageID string) (st *discordgo.Message, err error)
	ChannelMessageSend(channelID string, content string) (*discordgo.Message, error)
	ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error)
	ChannelMessageEditEmbed(channelID string, messageID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error)
	MessageReactionAdd(channelID string, messageID string, emojiID string) error
	MessageReactionsRemoveAll(channelID string, messageID string) error
	User(userID string) (st *discordgo.User, err error)
}

type SessDAL struct {
	ds discordgo.Session
}

func (s *SessDAL) Open() error {
	return s.ds.Open()
}

func (s *SessDAL) Close() error {
	return s.ds.Close()
}

func (s *SessDAL) AddHandler(handler interface{}) func() {
	return s.ds.AddHandler(handler)
}

func (s *SessDAL) ChannelMessage(channelID string, messageID string) (st *discordgo.Message, err error) {
	return s.ds.ChannelMessage(channelID, messageID)
}
func (s *SessDAL) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	return s.ds.ChannelMessageSend(channelID, content)
}

func (s *SessDAL) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return s.ds.ChannelMessageSendEmbed(channelID, embed)
}

func (s *SessDAL) ChannelMessageEditEmbed(channelID string, messageID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return s.ds.ChannelMessageEditEmbed(channelID, messageID, embed)
}

func (s *SessDAL) MessageReactionAdd(channelID string, messageID string, emojiID string) error {
	return s.ds.MessageReactionAdd(channelID, messageID, emojiID)
}

func (s *SessDAL) MessageReactionsRemoveAll(channelID string, messageID string) error {
	return s.ds.MessageReactionsRemoveAll(channelID, messageID)
}

func (s *SessDAL) User(userID string) (st *discordgo.User, err error) {
	return s.ds.User(userID)
}

func NewSessDAL(token string) (*SessDAL, error) {
	dg, err := discordgo.New(token)
	if err != nil {
		return nil, oops.Wrapf(err, "error creating discord SessDAL")
	}

	return &SessDAL{
		ds: *dg,
	}, nil
}
