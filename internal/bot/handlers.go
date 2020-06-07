package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/ecshreve/civ-bot-go/internal/util"
)

func (b *Bot) validateMessageHandlerArgs(m *discordgo.Message) Command {
	// Ignore all messages created by the bot itself.
	if m.Author.ID == b.DS.State.User.ID {
		return nil
	}

	// Ignore all messages that don't have the /civ prefix.
	if !strings.HasPrefix(m.Content, "/civ") {
		return nil
	}

	// Get the arguments.
	args := strings.Split(m.Content, " ")[1:]

	// Ensure command exists.
	if len(args) == 0 {
		b.DS.ChannelMessageSend(m.ChannelID, util.ErrorMessage("command missing", "for a list of commands type `/civ help`"))
		return nil
	}

	c, ok := b.CommandMap[CommandID(args[0])]
	if !ok {
		b.DS.ChannelMessageSend(m.ChannelID, util.ErrorMessage("invalid command", "for a list of commands type `/civ help`"))
		return nil
	}

	return c
}

// MessageHandler processes any new Messages in a channel where the Bot is a Member.
func (b *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	command := b.validateMessageHandlerArgs(m.Message)
	if command == nil {
		return
	}

	command.Process(b, m.Message)
}

// ReactionHandler processes any ReactionAdds in a channel where the Bot is a Member.
func (b *Bot) ReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	fmt.Println("reactionhandler")
}
