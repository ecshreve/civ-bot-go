package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/util"
)

// CommandsHandler handles all civ-bot commands.
func CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't have the /civ prefix.
	if !strings.HasPrefix(m.Content, "/civ") {
		return
	}

	// Get the arguments.
	args := strings.Split(m.Content, " ")[1:]

	// Ensure command exists.
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, util.ErrorMessage("command missing", "for a list of commands type `/civ help`"))
		return
	}

	// Call the corresponding handler.
	switch args[0] {
	case "help":
		helpCommandHandler(s, m, args)
	case "new", "oops":
		newCommandHandler(s, m)
	case "info":
		infoCommandHandler(s, m)
	case "list":
		listCommandHandler(s, m)
	case "ban":
		banCommandHandler(s, m, args)
	default:
		s.ChannelMessageSend(m.ChannelID, util.ErrorMessage("invalid command", "for a list of help topics, type `/civ help`"))
	}
}

// ReactionsHandler handles all civ-bot related reactions.
func ReactionsHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignore all reactions created by the bot itself.
	if r.UserID == s.State.User.ID {
		return
	}

	// Fetch some extra information about the message associated to the reaction.
	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)

	// Ignore reactions on messages that have an error or that have not been sent by the bot.
	if err != nil || m == nil || m.Author.ID != s.State.User.ID {
		return
	}

	// Ignore messages that are not embeds with a command in the footer.
	if len(m.Embeds) != 1 || m.Embeds[0].Footer == nil || m.Embeds[0].Footer.Text == "" {
		return
	}

	// Ignore reactions that haven't been set by the bot.
	if !util.IsBotReaction(s, m.Reactions, &r.Emoji) {
		return
	}

	// Ignore when sender is invalid or is a bot.
	user, err := s.User(r.UserID)
	if err != nil || user == nil || user.Bot {
		return
	}

	args := m.Embeds[0].Footer.Text

	// Call the corresponding handler.
	switch args {
	case "new":
		newReactionHandler(s, r, m, user)
	case "pick":
		pickReactionHandler(s, r, m, user)
	}
}
