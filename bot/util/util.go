package util

import "github.com/bwmarrin/discordgo"

// CivKey represents an integer key for a Civ.
type CivKey int

// IsBotReaction checks if users reaction is one preset by the bot.
func IsBotReaction(s *discordgo.Session, reactions []*discordgo.MessageReactions, emoji *discordgo.Emoji) bool {
	for _, r := range reactions {
		if r.Emoji.Name == emoji.Name && r.Me {
			return true
		}
	}

	return false
}
