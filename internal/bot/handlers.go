package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("commandhandler")
}

func (b *Bot) ReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	fmt.Println("reactionhandler")
}
