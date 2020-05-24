package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Formats the user in a readable format.
func formatUser(u *discordgo.User) string {
	return fmt.Sprintf("<@%s>", u.ID)
}

func formatCiv(c *Civ) string {
	return fmt.Sprintf("%s -- %s", c.CivBase, c.LeaderBase)
}

func formatCivs(cs []*Civ) string {
	ret := ""
	for _, c := range cs {
		ret = ret + formatCiv(c) + "\n"
	}
	return ret
}

// Generic message format for errors.
func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
