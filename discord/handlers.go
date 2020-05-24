package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandsHandler handles all civ-bot commands.
func CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "help":
		// Help command with topic
		if len(args) > 1 {
			helpCommandHandler(s, m, args[1])
		} else { // Help command without topic
			helpCommandHandler(s, m, "")
		}
	default:
		s.ChannelMessageSend(m.ChannelID, errorMessage("invalid command", "for a list of help topics, type `/civ help`"))
	}
}
