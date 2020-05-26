package util

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// IsBotReaction checks if users reaction is one preset by the bot.
func IsBotReaction(s *discordgo.Session, reactions []*discordgo.MessageReactions, emoji *discordgo.Emoji) bool {
	for _, r := range reactions {
		if r.Emoji.Name == emoji.Name && r.Me {
			return true
		}
	}

	return false
}

// FormatUser returns a string for a single user in a readable format.
func FormatUser(u *discordgo.User) string {
	return fmt.Sprintf("<@%s>", u.ID)
}

// FormatUsers returns a string for a slice of users in a readable format.
func FormatUsers(players map[string]*discordgo.User) string {
	ret := ""
	for _, p := range players {
		ret = ret + FormatUser(p) + "\n"
	}
	return ret
}

// ErrorMessage returns a string for a generic error message.
func ErrorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
