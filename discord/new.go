package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (cs *CivSession) newCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	title := "🆕 starting a new game"
	description := "- whoever wants to play react with  ✋\n- someone add a  ✅ react when ready to continue \n- enter `/civ oops` at any point to completely start over"

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

	// Reset the CivSession and add the two reactions needed to add players to the
	// game, and complete adding players to the game.
	cs.reset()
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✋")
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✅")
}

// newReactionHandler handles all new related reactions.
func (cs *CivSession) newReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "✋" {
		cs.Players[user.ID] = user
	}
	if r.Emoji.Name == "✅" {
		cs.banInstructions(s, m.ChannelID)
	}
}
