package discord

import "github.com/bwmarrin/discordgo"

// isBotReaction checks if users reaction is one preset by the bot.
func isBotReaction(s *discordgo.Session, reactions []*discordgo.MessageReactions, emoji *discordgo.Emoji) bool {
	for _, r := range reactions {
		if r.Emoji.Name == emoji.Name && r.Me {
			return true
		}
	}

	return false
}
