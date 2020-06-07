package bot

import (
	"log"
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

	readyToPick := false
	if _, ok := command.(*banCommand); ok == true {
		readyToPick = b.ReadyToPick()
	}

	if !readyToPick {
		return
	}

	b.CivState.DoRepick = true
	for b.CivState.DoRepick {
		err := b.Pick(m.ChannelID)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (b *Bot) validateReactionHandlerArgs(r *discordgo.MessageReaction) Reaction {
	// Ignore all reactions created by the bot itself.
	if r.UserID == b.DS.State.User.ID {
		return nil
	}

	// Fetch some extra information about the message associated to the reaction.
	m, err := b.DS.ChannelMessage(r.ChannelID, r.MessageID)

	// Ignore reactions on messages that have an error or that have not been sent by the bot.
	if err != nil || m == nil || m.Author.ID != b.DS.State.User.ID {
		return nil
	}

	// Ignore messages that are not embeds with a command in the footer.
	if len(m.Embeds) != 1 || m.Embeds[0].Footer == nil || m.Embeds[0].Footer.Text == "" {
		return nil
	}

	// Ignore reactions that haven't been set by the bot.
	if !util.IsBotReaction(b.DS.Session, m.Reactions, &r.Emoji) {
		return nil
	}

	// Ignore when sender is invalid or is a bot.
	user, err := b.DS.User(r.UserID)
	if err != nil || user == nil || user.Bot {
		return nil
	}

	args := m.Embeds[0].Footer.Text
	rr, ok := b.ReactionMap[CommandID(args)]
	if !ok {
		b.DS.ChannelMessageSend(m.ChannelID, util.ErrorMessage("invalid reaction", "this should never happen lol"))
		return nil
	}

	return rr
}

// ReactionHandler processes any ReactionAdds in a channel where the Bot is a Member.
func (b *Bot) ReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	reaction := b.validateReactionHandlerArgs(r.MessageReaction)
	if reaction == nil {
		return
	}

	reaction.Process(b, r.MessageReaction)
}
