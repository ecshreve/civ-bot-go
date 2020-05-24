package discord

import (
	"fmt"

	"github.com/kr/pretty"

	"github.com/bwmarrin/discordgo"
)

func newCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, cs *CivSession) {
	title := "🆕 starting a new game"
	description := "whoever wants to play react with  ✋\n someone add a  ✅ react when ready to continue"

	newSession, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       cDARKPURPLE,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "new",
		},
	})

	if err != nil {
		fmt.Println("error creating new session")
		return
	}

	cs.reset()
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✋")
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✅")
}

// newReactionHandler handles all new related reactions.
func newReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, cs *CivSession, user *discordgo.User) {
	if r.Emoji.Name == "✋" {
		cs.Players = append(cs.Players, user)
	}
	if r.Emoji.Name == "✅" {
		pretty.Println(cs)
	}
}
