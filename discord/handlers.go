package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandsHandler handles all civ-bot commands.
func (cs *CivSession) CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't have the !checkers prefix
	if !strings.HasPrefix(m.Content, "/civ") {
		return
	}

	// Get the arguments
	args := strings.Split(m.Content, " ")[1:]
	// Ensure valid command
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("command missing", "for a list of commands type `/civ help`"))
		return
	}

	// Call the corresponding handler
	switch args[0] {
	case "help":
		// Help command with topic
		if len(args) > 1 {
			helpCommandHandler(s, m, args[1])
		} else { // Help command without topic
			helpCommandHandler(s, m, "")
		}
	case "new", "oops":
		newCommandHandler(s, m, cs)
	case "info":
		infoCommandHandler(s, m, cs)
	case "list":
		listCommandHandler(s, m, cs)
	default:
		s.ChannelMessageSend(m.ChannelID, errorMessage("invalid command", "for a list of help topics, type `/civ help`"))
	}
}

// ReactionsHandler handles all checkers related reactions
func (cs *CivSession) ReactionsHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignore all reactions created by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	// Fetch some extra information about the message associated to the reaction
	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	// Ignore reactions on messages that have an error or that have not been sent by the bot
	if err != nil || m == nil || m.Author.ID != s.State.User.ID {
		return
	}

	// Ignore messages that are not embeds with a command in the footer
	if len(m.Embeds) != 1 || m.Embeds[0].Footer == nil || m.Embeds[0].Footer.Text == "" {
		return
	}

	// Ignore reactions that haven't been set by the bot
	if !isBotReaction(s, m.Reactions, &r.Emoji) {
		return
	}

	user, err := s.User(r.UserID)
	// Ignore when sender is invalid or is a bot
	if err != nil || user == nil || user.Bot {
		return
	}

	args := m.Embeds[0].Footer.Text

	// Call the corresponding handler
	switch args {
	case "new":
		newReactionHandler(s, r, m, cs, user)
	}
}
