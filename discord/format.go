package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// formatUser returns a string for a single user in a readable format.
func formatUser(u *discordgo.User) string {
	return fmt.Sprintf("<@%s>", u.ID)
}

// formatUsers returns a string for a slice of users in a readable format.
func formatUsers(players map[string]*discordgo.User) string {
	ret := ""
	for _, p := range players {
		ret = ret + formatUser(p) + "\n"
	}
	return ret
}

// formatCiv returns a string for a single Civ in a readable format.
func formatCiv(c *Civ) string {
	formatStr := "[%s -- %s](%s)"
	return fmt.Sprintf(formatStr, c.CivBase, c.LeaderBase, c.ZigURL)
}

// formatCivs returns a string for a slice of Civs in a readable format.
//
// TODO: fix this, just call it on CivSession pointer receiver.
func formatCivs(cs []*Civ) string {
	ret := ""
	for _, c := range cs {
		ret = ret + "\n" + formatCiv(c)
	}
	return ret
}

// formatPicks returns a string in a readable format for each player's picks.
func formatPicks(picks map[*discordgo.User][]*Civ) string {
	ret := ""
	for k, v := range picks {
		ret = ret + "\n" + formatUser(k) + ":\n" + formatCivs(v) + "\n-----\n"
	}
	return ret
}

// formatBans returns a string for the CivSession's Bans in a readable format.
func (cs *CivSession) formatBans() string {
	bans := cs.Bans
	if bans == nil || len(bans) == 0 {
		return "no bans yet"
	}
	ret := ""
	for k, v := range bans {
		ret = ret + formatUser(cs.Players[k]) + ": " + formatCiv(v) + "\n"
	}
	return ret
}

// errorMessage returns a string for a generic error message.
func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
