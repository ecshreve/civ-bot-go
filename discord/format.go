package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Formats the user in a readable format.
func formatUser(u *discordgo.User) string {
	return fmt.Sprintf("<@%s>", u.ID)
}

// Generic message format for errors.
func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
