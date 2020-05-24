package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Formats the user in a readable format.
func formatUser(u *discordgo.User) string {
	return fmt.Sprintf("<@%s>", u.ID)
}

func formatUsers(players map[string]*discordgo.User) string {
	ret := ""
	for _, p := range players {
		ret = ret + formatUser(p) + "\n"
	}
	return ret
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

func formatPicks(picks map[*discordgo.User][]*Civ) string {
	ret := ""
	for k, v := range picks {
		ret = ret + formatUser(k) + ":\n" + formatCivs(v) + "\n--\n\n"
	}
	return ret
}

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

// Generic message format for errors.
func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
