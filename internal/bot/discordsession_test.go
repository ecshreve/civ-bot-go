package bot

import "github.com/bwmarrin/discordgo"

func (s *DiscordSession) Open() error {
	mock.Input(interface{}(s.Open))
	return nil
}

func (s *DiscordSession) Close() error {
	mock.Input(interface{}(s.Close))
	return nil
}

func (s *DiscordSession) AddHandler(handler interface{}) func() {
	mock.Input(interface{}(s.AddHandler), handler)
	return func() {}
}

func (s *DiscordSession) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	mock.Input(interface{}(s.ChannelMessageSend), channelID, content)
	return &discordgo.Message{
		ChannelID: channelID,
		Content:   content,
	}, nil
}

func (s *DiscordSession) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	mock.Input(interface{}(s.ChannelMessageSendEmbed), channelID, embed)
	return &discordgo.Message{
		ChannelID: channelID,
		Embeds:    []*discordgo.MessageEmbed{embed},
	}, nil
}

func (s *DiscordSession) MessageReactionAdd(channelID string, messageID string, emojiID string) error {
	mock.Input(interface{}(s.MessageReactionAdd), channelID, messageID, emojiID)
	return nil
}

func (s *DiscordSession) User(userID string) (*discordgo.User, error) {
	mock.Input(interface{}(s.User), userID)
	return &discordgo.User{
		ID: userID,
	}, nil
}
